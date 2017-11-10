package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/k0kubun/pp"
	"github.com/redite/kleng/core"
	"github.com/redite/kleng/utils"
)

func handleNoWWW(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	http.Redirect(w, r, "http://"+vars["domain"]+"."+vars["tld"], 301)
}

func listStared(w http.ResponseWriter, r *http.Request) {
	c := core.NewGithubClient()

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
