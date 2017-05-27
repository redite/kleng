package core

import (
	"github.com/google/go-github/github"
	"github.com/redite/kleng/config"
	"golang.org/x/oauth2"
)

// GithubClient wrap go-github
type GithubClient struct {
	C        *github.Client
	Username string
}

type StarredRepo struct {
	Fullname string
}

// NewGithubClient return client for github
func NewGithubClient() GithubClient {

	c, _ := config.LoadConfig()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: c.Token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	g := GithubClient{
		C:        github.NewClient(tc),
		Username: c.Username,
	}
	return g
}

// ListStarred return list of Username's stared repository
func (g GithubClient) ListStarred(page int) ([]*github.StarredRepository, error) {
	opt := new(github.ActivityListStarredOptions)
	opt.Sort = "created"
	opt.Direction = "ascending"
	opt.Page = page

	stared, _, err := g.C.Activity.ListStarred(g.Username, opt)
	if err != nil {
		return nil, err
	}

	return stared, err
}

func (g GithubClient) PopulateStarred() error {
	// for {
	/// todo
	// }

	return nil
}
