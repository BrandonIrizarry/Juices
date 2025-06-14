package main

import (
	"log"
	"net/http"
	"strconv"
)

var idsToCounts = make(map[string]int)

func postCount(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving %s %s\n", r.Method, r.URL.Path)

	// Get the count.
	count, err := strconv.Atoi(r.FormValue("count"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Got count: %d\n", count)

	// Record the count under the given ID (we'll accumulate
	// counts under each date later.)
	id := r.Header.Get("Hx-Trigger")

	if id == "" {
		message := "This hypermedia element requires an id attribute"
		log.Println(message)
		http.Error(w, message, http.StatusInternalServerError)
	}

	idsToCounts[id] = count

	log.Printf("Current counts: %v\n", idsToCounts)
}
