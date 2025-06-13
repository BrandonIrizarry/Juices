package main

import (
	"log"
	"net/http"

	"github.com/BrandonIrizarry/juices/internal/juicedate"
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

	// Endpoints (each handler is defined in its own file inside
	// the main package.
	mux.HandleFunc("GET /date", juicedate.GetDate)
	mux.HandleFunc("POST /date", juicedate.PostDate)
	mux.HandleFunc("DELETE /date", deleteDate)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Fatal(srv.ListenAndServe())
}
