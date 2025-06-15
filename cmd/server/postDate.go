package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var index int

// PostDate serves a row consisting of the selected date, an HTML5
// counter widget, and a Delete button. If this is an Add Date
// operation, another Add Date button is appended to the served HTML.
func postDate(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving %s %s\n", r.Method, r.URL.Path)

	w.Header().Set("Content-Type", "text/html")

	date, err := parseDate(r.FormValue("date"))

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	itemName := r.PathValue("itemName")

	if itemName == "" {
		err := errors.New("Missing 'itemName' path value")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	entryHTML, err := template.New("entry").Funcs(template.FuncMap{
		"inc": func() int {
			index++
			return index
		},
	}).ParseFiles("assets/entry.html")

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type dataView struct {
		ItemName string
		Date     string
	}

	dv := dataView{itemName, date}

	if err := entryHTML.ExecuteTemplate(w, "entry", dv); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func parseDate(dateFormValue string) (string, error) {
	parts := strings.SplitN(dateFormValue, "-", 2)

	if len(parts) != 2 {
		return "", fmt.Errorf("Invalid date: %s", dateFormValue)
	}

	return parts[1], nil
}
