package main

import (
	"log"
	"net/http"

	"github.com/BrandonIrizarry/juices/internal/cid"
	"github.com/BrandonIrizarry/juices/internal/juicecount"
)

func deleteDate(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving %s %s\n", r.Method, r.URL.Path)

	canonicalID, err := cid.ParseCanonicalID(r.Header.Get("Hx-Trigger"))

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	juicecount.Delete(canonicalID)

	log.Println(juicecount.Info())
}
