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

	// Serve page as a static asset.
	serveMainPage := http.StripPrefix("/app/", http.FileServer(http.Dir("./app")))

	// Static stuff.
	mux.Handle("GET /js/", loadJS)
	mux.Handle("GET /app/", serveMainPage)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Fatal(srv.ListenAndServe())
}
