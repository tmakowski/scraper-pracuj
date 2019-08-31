package server

import (
	"log"
	"net/http"
)

// Function starts the server on provided port and opens a default browser.
func Start(port int) {
	// Setting up endpoints and server
	c := Client{port}
	http.HandleFunc("/", c.mainHandler)
	http.HandleFunc("/pracuj", c.pracujHandler)

	// Printing start-up info
	log.Printf("[Version: %s] Started logging...\n", CurrentVersion().toString())
	log.Println("Listening on", c.url())

	// Starting browser
	err := c.startBrowser()
	if err != nil {
		log.Println("[ERROR] Could not open browser. Please visit:", c.url())
	}

	// Starting the server
	log.Fatal(c.startServer())
}
