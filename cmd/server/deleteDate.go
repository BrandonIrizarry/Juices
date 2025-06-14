package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/BrandonIrizarry/juices/internal/juicecount"
)

func deleteDate(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving %s %s\n", r.Method, r.URL.Path)

	id := r.Header.Get("Hx-Trigger")

	if id == "" {
		message := "Delete button is missing ID attribute"
		log.Println(message)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}

	indexStr, found := strings.CutPrefix(id, "delete-")

	if !found {
		message := fmt.Sprintf("Invalid delete-button id: %s", id)
		log.Println(message)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}

	deleteButtonIndex, err := strconv.Atoi(indexStr)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var foundDateID string

	for dateID := range juicecount.IDsToCounts {
		parts := strings.SplitN(dateID, "-", 3)

		if len(parts) != 3 {
			message := fmt.Sprintf("Invalid map-key/counter ID: %s", dateID)
			log.Println(message)
			http.Error(w, message, http.StatusInternalServerError)
			return
		}

		dateIDIndex, err := strconv.Atoi(parts[2])

		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if dateIDIndex == deleteButtonIndex {
			foundDateID = dateID
			break
		}
	}

	// FIXME: normally, a row should have a count - except
	// whenever the counter hasn't been used yet. Could we use a
	// default value attribute for the counter, then redirect POST
	// /date to POST /count? In that case, here would then guard
	// against foundDateID staying empty.
	if foundDateID == "" {
		log.Println("Warning: counter is either uninitialized, or an error has occured.")
	}

	delete(juicecount.IDsToCounts, foundDateID)

	log.Printf("Current counts: %v\n", juicecount.IDsToCounts)
}
