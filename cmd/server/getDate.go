package main

import (
	"errors"
	"html/template"
	"log"
	"net/http"
)

// GetDate serves the HTML5 date widget to the page.
func getDate(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving %s %s\n", r.Method, r.URL.Path)

	w.Header().Set("Content-Type", "text/html")

	itemName := r.PathValue("itemName")

	if itemName == "" {
		err := errors.New("Missing 'itemName' path value")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	datePickerHTML, err := template.New("datepicker").ParseFiles("assets/datepicker.html")

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := datePickerHTML.ExecuteTemplate(w, "datepicker", itemName); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
