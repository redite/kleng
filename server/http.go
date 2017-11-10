package server

import (
	"net/http"

	"fmt"

	"github.com/gorilla/mux"
	"github.com/redite/kleng/config"
)

// RunServer ...
func RunServer(c config.Config) error {
	fmt.Printf("creating server at port %d\n", c.Port)

	router := mux.NewRouter()

	// Force no WWW
	router.HandleFunc("/", handleNoWWW).Host("www.{domain}.{tld}")

	// Serve API
	router.HandleFunc("/starred", listStared)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

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
