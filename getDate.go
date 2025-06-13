package main

import (
	"log"
	"net/http"
	"strings"
)

var html = strings.TrimSpace(`
<input type="date"/>
`)

func getDate(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving %s %s\n", r.Method, r.URL.Path)

	w.Header().Set("Content-Type", "text/html")

	w.Write([]byte(html))
}
