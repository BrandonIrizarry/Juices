package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

var getDateHTML = strings.TrimSpace(`
<input type="date" name="date" hx-post="/date" hx-swap="outerHTML"/>
`)

var postDateHTML = strings.TrimSpace(`
<div>
<span>%s</span>
<button id="edit" hx-get="/date" hx-swap="outerHTML" hx-target="closest div">Edit</span>
</div>
%s`)

var addDateButton = strings.TrimSpace(`
<button id="add" hx-get="/date" hx-swap="outerHTML">Add Date</button>
`)

var hxTrigger string

func getDate(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving %s %s\n", r.Method, r.URL.Path)

	w.Header().Set("Content-Type", "text/html")

	hxTrigger = r.Header.Get("Hx-Trigger")

	log.Printf("Element: %s\n", hxTrigger)

	_, err := w.Write([]byte(getDateHTML))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func postDate(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving %s %s\n", r.Method, r.URL.Path)

	w.Header().Set("Content-Type", "text/html")

	date := strings.SplitN(r.FormValue("date"), "-", 2)[1]

	finalHTML, err := computeDateFinalHTML(date)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte(finalHTML))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func computeDateFinalHTML(date string) (string, error) {
	switch hxTrigger {
	case "edit":
		return fmt.Sprintf(postDateHTML, date, ""), nil
	case "add":
		return fmt.Sprintf(postDateHTML, date, addDateButton), nil
	default:
		return "", fmt.Errorf("Invalid HX trigger: %s", hxTrigger)
	}
}
