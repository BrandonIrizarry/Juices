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
<input type="number" name="count" hx-post="/count" hx-trigger="change delay:1s" min=0 id="%[1]s-%d" />
<button id="delete" hx-delete="/date" hx-swap="delete" hx-target="closest div">Delete</button>
</div>
%s`)

var addDateButton = strings.TrimSpace(`
<button name="add" hx-get="/date" hx-swap="outerHTML">Add Date</button>
`)

var hxTriggerName string

var countElementID = 0

func GetDate(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving %s %s\n", r.Method, r.URL.Path)

	w.Header().Set("Content-Type", "text/html")

	hxTriggerName = r.Header.Get("Hx-Trigger-Name")

	log.Printf("ID of triggering element (add/edit): %s\n", hxTriggerName)

	if hxTriggerName != "add" && hxTriggerName != "edit" {
		message := fmt.Sprintf("Unexpected trigger element: %s", hxTriggerName)
		log.Println(message)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}

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
		countElementID++
		return fmt.Sprintf(postDateHTML, date, countElementID, ""), nil
	case "add":
		countElementID++
		return fmt.Sprintf(postDateHTML, date, countElementID, addDateButton), nil
	default:
		return "", fmt.Errorf("Invalid HX trigger name: %s", hxTriggerName)
	}
}
