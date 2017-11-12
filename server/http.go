package server

import (
	"net/http"
	"os"

	"fmt"

	"github.com/gorilla/mux"
	"github.com/redite/kleng/config"
)

func initRoutes(mx *mux.Router) {
	webRoot, err := os.Getwd()
	if err != nil {
		panic("could not retrieve working directory")
	}

	fmt.Println(webRoot)

	mx.PathPrefix("/").Handler(http.FileServer(http.Dir(webRoot + "/public/")))
}

// RunServer ...
func RunServer(c config.Config) error {
	fmt.Printf("creating server at port %d\n", c.Port)

	router := mux.NewRouter()

	// Force no WWW
	router.HandleFunc("/", handleNoWWW).Host("www.{domain}.{tld}")

	// Serve API
	router.HandleFunc("/api/starred", listStared)

	initRoutes(router)

	// Set the router
	http.Handle("/", router)

	// Start the server
	fmt.Println("Listening...")
	err := http.ListenAndServe(
		fmt.Sprintf(":%d", c.Port),
		router,
	)

	return err
}
