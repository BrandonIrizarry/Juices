package main

import (
	"log"
	"net/http"
)

func main() {
	// Set up server.
	mux := http.NewServeMux()

	// For now, we need this primarily to load HTMX.
	loadJS := http.StripPrefix("/js/", http.FileServer(http.Dir("./js")))
	mux.Handle("GET /js/", loadJS)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Fatal(srv.ListenAndServe())
}
