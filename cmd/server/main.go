package main

import (
	"log"
	"net/http"

	"github.com/BrandonIrizarry/juices/internal/juicedate"
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
	mux.HandleFunc("GET /date", juicedate.GetDate)
	mux.HandleFunc("POST /date", juicedate.PostDate)
	mux.HandleFunc("DELETE /date", deleteDate)
	mux.HandleFunc("POST /count", postCount)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Fatal(srv.ListenAndServe())
}
