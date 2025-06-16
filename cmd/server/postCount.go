package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type entry struct {
	itemName string
	date     string
	count    int
}

// Map counter ID attributes to entry info.
var counts = make(map[string]entry)

func postCount(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving %s %s\n", r.Method, r.URL.Path)

	// Get the count.
	count, err := strconv.Atoi(r.FormValue("count"))

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	itemName, err := nonEmptyValue(r.PathValue("itemName"))

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	date, err := nonEmptyValue(r.PathValue("date"))

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	e := entry{itemName, date, count}

	// Use this counter's ID attribute as the map key.
	id, err := nonEmptyValue(r.Header.Get("Hx-Trigger"))

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	counts[id] = e

	log.Printf("Counts: %v\n", counts)
}

func nonEmptyValue(value string) (string, error) {
	if value == "" {
		return "", fmt.Errorf("Missing '%s' from path values", value)
	}

	return value, nil
}
