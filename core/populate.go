package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"strconv"

	"github.com/boltdb/bolt"
	"github.com/google/go-github/github"
	"github.com/redite/kleng/utils"
)

type requestControl struct {
	starred []*github.StarredRepository
	err     []error
	wg      sync.WaitGroup
	mu      *sync.RWMutex
}

// Populate perform call to user's starred repo on github and store the result to bolt database
func Populate() error {
	return firstRun()
}

func firstRun() error {

	rc := &requestControl{
		mu: new(sync.RWMutex),
	}

	c := utils.NewGithubClient()
	pageCount, _ := getStarredPageCount(c.Username)

	// diparallel aja biar cepet.
	for page := 0; page <= pageCount; page++ {
		rc.wg.Add(1)
		go func(idx int, stack *requestControl) {
			defer stack.wg.Done()

			res, err := c.ListStarred(idx)
			if err != nil {
				stack.mu.Lock()
				stack.err = append(stack.err, err)
				stack.mu.Unlock()
				return
			}

			// ga boleh rebutan, nanti dimarahin mama.
			stack.mu.Lock()
			stack.starred = append(stack.starred, res...)
			stack.mu.Unlock()
		}(page, rc)
	}

	// tungguin sampe goroutine kelar semua.
	rc.wg.Wait()

	if len(rc.err) > 0 {
		fmt.Printf("got %d error", len(rc.err))
		fmt.Println(rc.err)
	}

	db, err := OpenDB()
	if err != nil {
		err = fmt.Errorf("failed opening DB: %s", err.Error())
		return err
	}

	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("kleng"))
		if err != nil {
			return err
		}

		j, err := json.Marshal(rc.starred)
		if err != nil {
			return fmt.Errorf("cannot marshal data: %s", err.Error())
		}

		err = bucket.Put([]byte("starred_repo"), j)
		return err
	})

	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("kleng"))
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found!", []byte("kleng"))
		}

		val := bucket.Get([]byte("starred_repo"))
		fmt.Println(string(val))

		return nil
	})

	// j, err := json.Marshal(rc.starred)
	// err = ioutil.WriteFile("./new-starred-1.json", j, 0644)
	// if err != nil {
	// return err
	// }

	return nil
}

func populationUpdate() error {
	return nil
}

func getStarredPageCount(username string) (count int, err error) {
	client := http.DefaultClient
	resp, err := client.Get(
		fmt.Sprintf("%s/%s/%s", "https://api.github.com/users", username, "starred"),
	)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	links := resp.Header.Get("Link")
	if len(links) > 0 {
		lastpage := strings.Replace(
			strings.Replace(
				strings.Split(links, ";")[1],
				"rel=\"next\", <",
				"",
				-1,
			),
			">",
			"",
			-1,
		)

		lasturl, err := url.Parse(lastpage)
		if err != nil {
			return count, err
		}

		count, err = strconv.Atoi(lasturl.Query().Get("page"))
	}

	return count, err
}
