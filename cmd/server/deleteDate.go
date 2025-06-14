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

	if err := juicecount.Delete(deleteButtonIndex); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println(juicecount.Info())
}
