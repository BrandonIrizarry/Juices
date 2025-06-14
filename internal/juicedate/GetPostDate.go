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
<button id="delete-%[2]d" hx-delete="/date" hx-swap="delete" hx-target="closest div" hx-confirm="Delete this row?">Delete</button>
</div>
%s`)

var addDateButton = strings.TrimSpace(`
<button name="add" hx-get="/date" hx-swap="outerHTML">Add Date</button>
`)

// The name of the widget requesting GET /date (could be "add" or "edit".)
var hxTriggerName string

// The index appended to the computed ID of a counter widget (input
// type=number.) This is used to enforce unique HTML id attribute
// values among the counter widgets themselves.
var countElementIndex = 0

// GetDate serves the HTML5 date widget to the page.
func GetDate(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving %s %s\n", r.Method, r.URL.Path)

	w.Header().Set("Content-Type", "text/html")

	hxTriggerName = r.Header.Get("Hx-Trigger-Name")

	if hxTriggerName == "" {
		message := "This hypermedia element requires a name attribute"
		log.Println(message)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}

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

// PostDate serves a row consisting of the selected date, an HTML5
// counter widget, and a Delete button. If this is an Add Date
// operation, another Add Date button is appended to the served HTML.
func PostDate(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving %s %s\n", r.Method, r.URL.Path)

	w.Header().Set("Content-Type", "text/html")

	submittedDate := r.FormValue("date")
	parts := strings.SplitN(submittedDate, "-", 2)

	if len(parts) != 2 {
		message := fmt.Sprintf("Invalid date: %s", submittedDate)
		log.Println(message)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}

	date := parts[1]

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

// computeDateFinalHTML returns the HTML to serve upon setting a date,
// dependent on whether the action (hxTriggerName) is an edit or an
// add.
func computeDateFinalHTML(date string) (string, error) {
	switch hxTriggerName {
	case "edit":
		countElementIndex++
		return fmt.Sprintf(postDateHTML, date, countElementIndex, ""), nil
	case "add":
		countElementIndex++
		return fmt.Sprintf(postDateHTML, date, countElementIndex, addDateButton), nil
	default:
		return "", fmt.Errorf("Invalid HX trigger name: %s", hxTriggerName)
	}
}
