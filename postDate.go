package main

import (
	"log"
	"net/http"
	"strings"
)

var postDateHTML = strings.TrimSpace(`
<span>Got date.</span>
<button hx-get="/date" hx-swap="outerHTML">Add Date</button>
`)

func postDate(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving %s %s\n", r.Method, r.URL.Path)

	w.Header().Set("Content-Type", "text/html")

	w.Write([]byte(postDateHTML))
}
