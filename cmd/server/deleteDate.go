package main

import (
	"log"
	"net/http"
	"strconv"
)

func deleteDate(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving %s %s\n", r.Method, r.URL.Path)

	// The key of the map entry we're deleting has precisely an id
	// of one less than that of this delete button.
	id, err := strconv.Atoi(r.Header.Get("Hx-Trigger"))

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	delete(counts, id-1)

	log.Printf("Counts: %v\n", counts)
}
