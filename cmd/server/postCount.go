package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/BrandonIrizarry/juices/internal/cid"
	"github.com/BrandonIrizarry/juices/internal/juicecount"
)

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
	canonicalID, err := cid.ParseCanonicalID(r.Header.Get("Hx-Trigger"))

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	juicecount.SetBulk(canonicalID, count)

	log.Println(juicecount.Info())
}
