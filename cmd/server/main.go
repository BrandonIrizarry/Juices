package main

import (
	"html/template"
	"log"
	"net/http"
)

var entryWithIndex func() (*template.Template, error)

func main() {
	// Set up server.
	mux := http.NewServeMux()

	// Serve page as a static asset (includes CSS and JS, where
	// HTMX resides)
	serveMainPage := http.StripPrefix("/app/", http.FileServer(http.Dir("./app")))
	mux.Handle("GET /app/", serveMainPage)
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/app", http.StatusSeeOther)
	})

	// Endpoints (each handler is defined in its own file inside
	// the main package.
	mux.HandleFunc("GET /date/{itemName}", getDate)
	mux.HandleFunc("POST /date/{itemName}", postDate)
	mux.HandleFunc("DELETE /date", deleteDate)
	mux.HandleFunc("POST /count/{itemName}/{date}", postCount)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Fatal(srv.ListenAndServe())
}

func initEntryWithIndex() func() (*template.Template, error) {
	var index int

	return func() (*template.Template, error) {
		entryHTML, err := template.New("entry").Funcs(template.FuncMap{
			"inc": func() int {
				index++
				return index
			},
		}).ParseFiles("assets/entry.html")

		if err != nil {
			return nil, err
		}

		return entryHTML, nil
	}
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	entryWithIndex = initEntryWithIndex()
}
