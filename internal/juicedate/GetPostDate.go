package juicedate

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
<div class="entry">
<span name="edit" hx-get="/date" hx-swap="outerHTML" hx-target="closest div">%s</span>
<input type="number" name="count" min=0 />
<button id="delete" hx-delete="/date" hx-swap="delete" hx-target="closest div">Delete</button>
</div>
%s`)

var addDateButton = strings.TrimSpace(`
<button name="add" hx-get="/date" hx-swap="outerHTML">Add Date</button>
`)

var hxTriggerName string

func GetDate(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving %s %s\n", r.Method, r.URL.Path)

	w.Header().Set("Content-Type", "text/html")

	hxTriggerName = r.Header.Get("Hx-Trigger-Name")

	log.Printf("Element: %s\n", hxTriggerName)

	_, err := w.Write([]byte(getDateHTML))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func PostDate(w http.ResponseWriter, r *http.Request) {
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
	switch hxTriggerName {
	case "edit":
		return fmt.Sprintf(postDateHTML, date, ""), nil
	case "add":
		return fmt.Sprintf(postDateHTML, date, addDateButton), nil
	default:
		return "", fmt.Errorf("Invalid HX trigger name: %s", hxTriggerName)
	}
}
