package main

import (
	"log"
	"net/http"

	api "github.com/yuucu/cursor_go_tutorial/handlers"
)

func main() {
	// Create hello handler instance
	helloHandler := api.NewHelloHandler()

	// Create server with the handler
	srv, err := api.NewServer(helloHandler)
	if err != nil {
		log.Fatal(err)
	}

	// Start server on port 1323
	log.Println("Starting server on :1323")
	if err := http.ListenAndServe(":1323", srv); err != nil {
		log.Fatal(err)
	}
}
