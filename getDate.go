package main

import (
	"log"
	"net/http"
	"strings"
)

var getDateHTML = strings.TrimSpace(`
<input type="date" name="date" hx-post="/date" hx-swap="outerHTML"/>
`)

func getDate(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving %s %s\n", r.Method, r.URL.Path)

	w.Header().Set("Content-Type", "text/html")

	log.Printf("Element: %s\n", r.Header.Get("Hx-Trigger"))

	_, err := w.Write([]byte(getDateHTML))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
