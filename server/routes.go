package server

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/k0kubun/pp"
	"github.com/redite/kleng/utils"
)

func handleNoWWW(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	http.Redirect(w, r, "http://"+vars["domain"]+"."+vars["tld"], 301)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	file, e := ioutil.ReadFile("./mix-manifest.json")
	if e != nil {
		os.Exit(1)
	}

	var m MixManifest
	json.Unmarshal(file, &m)

	t, _ := template.ParseFiles("./resources/views/index.html")
	t.Execute(w, m)
}

func listStared(w http.ResponseWriter, r *http.Request) {
	c := utils.NewGithubClient()

	starred, err := c.ListStarred(0)
	if err != nil {
		errr := utils.Error{
			Status: 400,
			Title:  err.Error(),
		}
		utils.WriteError(errr, w)
		return
	}

	for _, s := range starred {
		pp.Println(s)
	}
}
