package server

import (
	"net/http"

	"fmt"

	"github.com/gorilla/mux"
	"github.com/redite/kleng/config"
)

type MixManifest struct {
	JS  string `json:"/dist/js/index.js"`
	CSS string `json:"/dist/css/app.css"`
}

// RunServer ...
func RunServer(c config.Config) error {
	fmt.Printf("creating server at port %d\n", c.Port)

	router := mux.NewRouter()

	// Force no WWW
	router.HandleFunc("/", handleNoWWW).Host("www.{domain}.{tld}")

	// Serve index
	router.HandleFunc("/", handleIndex)

	// Server assets
	router.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", http.FileServer(http.Dir("./dist/"))))
	router.PathPrefix("/fonts/").Handler(http.StripPrefix("/fonts/", http.FileServer(http.Dir("./fonts/"))))

	// Serve API
	router.HandleFunc("/starred", listStared)

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
