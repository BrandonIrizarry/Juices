package main

import (
	"log"
	"net/http"
)

func postClear(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving %s %s\n", r.Method, r.URL.Path)

	clear(counts)
}
