package main

import (
	"log"
	"net/http"
)

var port = ":8080"

func main() {
	srv := &http.Server{
		Addr:    port,
		Handler: routes(),
	}
	log.Println("Listening on port 8080")
	err := srv.ListenAndServe()
	log.Fatal(err)
}
