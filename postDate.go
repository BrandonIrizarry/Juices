package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

var postDateHTML = strings.TrimSpace(`
<span>%s</span>
<button hx-get="/date" hx-swap="outerHTML">Add Date</button>
`)

func postDate(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving %s %s\n", r.Method, r.URL.Path)

	w.Header().Set("Content-Type", "text/html")

	date := strings.SplitN(r.FormValue("date"), "-", 2)[1]

	w.Write([]byte(fmt.Sprintf(postDateHTML, date)))
}
