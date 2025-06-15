package main

import (
	"log"
	"net/http"
)

func main() {
	// Set up server.
	mux := http.NewServeMux()

	// Serve page as a static asset (includes CSS and JS, where
	// HTMX resides)
	serveMainPage := http.StripPrefix("/app/", http.FileServer(http.Dir("./app")))
	mux.Handle("GET /app/", serveMainPage)

	// Endpoints (each handler is defined in its own file inside
	// the main package.
	mux.HandleFunc("GET /date/{itemName}", getDate)
	mux.HandleFunc("POST /date/{itemName}", postDate)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Fatal(srv.ListenAndServe())
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
