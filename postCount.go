package main

import (
	"log"
	"net/http"
	"strconv"
)

func postCount(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving %s %s\n", r.Method, r.URL.Path)

	count, err := strconv.Atoi(r.FormValue("count"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Got count: %d\n", count)
}
