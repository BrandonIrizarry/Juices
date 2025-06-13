package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

var postDateHTML = strings.TrimSpace(`
<div>
<span>%s</span>
<button id="edit" hx-get="/date" hx-swap="outerHTML" hx-target="closest div">Edit</span>
</div>
<button id="add" hx-get="/date" hx-swap="outerHTML">Add Date</button>
`)

func postDate(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving %s %s\n", r.Method, r.URL.Path)

	w.Header().Set("Content-Type", "text/html")

	date := strings.SplitN(r.FormValue("date"), "-", 2)[1]

	_, err := w.Write([]byte(fmt.Sprintf(postDateHTML, date)))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
