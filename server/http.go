package server

import (
	"net/http"

	"fmt"

	"github.com/gorilla/mux"
	"github.com/redite/kleng/utils"
)

// RunServer ...
func RunServer(c utils.Config) error {
	fmt.Printf("creating server at port %d\n", c.Port)

	router := mux.NewRouter()

	// Force no WWW
	router.HandleFunc("/", handleNoWWW).Host("www.{domain}.{tld}")

	// Serve index
	router.HandleFunc("/", handleIndex)

	// Serve API
	router.HandleFunc("/starred", listStared)

	// Serve CSS
	router.HandleFunc("/css/style.css", handleStatic)

	// Serve JavaScript
	router.HandleFunc("/js/react-dom.min.js", handleStatic)
	router.HandleFunc("/js/react-with-addons.min.js", handleStatic)
	router.HandleFunc("/js/react-with-addons.js", handleStatic)
	router.HandleFunc("/js/resounden.js", handleStatic)
	router.HandleFunc("/js/resounden-utils.js", handleStatic)

	// Handle not found
	router.NotFoundHandler = http.HandlerFunc(handleNotFound)

	// Set the router
	http.Handle("/", router)

	// Start the server
	fmt.Println("Listening...")
	err := http.ListenAndServe(
		fmt.Sprintf(":%d", c.Port),
		nil,
	)

	return err
}
