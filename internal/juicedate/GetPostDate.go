package juicedate

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/BrandonIrizarry/juices/internal/juicecount"
)

var getDateHTML = strings.TrimSpace(`
<input type="date" name="date" hx-post="/date" hx-swap="outerHTML"/>
`)

func createSpan(index int, date string) string {
	spanHTML := strings.TrimSpace(`
<span name="edit" id="edit-%d" hx-get="/date" hx-swap="outerHTML" hx-target="closest div">%s</span>
`)

	return fmt.Sprintf(spanHTML, index, date)
}

func createCounter(index int, date string) string {
	counterHTML := strings.TrimSpace(`
<input type="number" name="count" hx-post="/count" hx-trigger="change delay:1s" min=0 id="%s-%d" value="0" />
`)
	return fmt.Sprintf(counterHTML, date, index)
}

func createDeleteButton(index int) string {
	deleteButtonHTML := strings.TrimSpace(`
<button id="delete-%d" hx-delete="/date" hx-swap="delete" hx-target="closest div" hx-confirm="Delete this row?">Delete</button>
`)
	return fmt.Sprintf(deleteButtonHTML, index)
}

func createEntry(index int, date string, wasAdd bool) string {
	entryHTML := strings.TrimSpace(`
<div class="entry">
%s
%s
%s
</div>
%s`)

	var addDateButton string

	if wasAdd {
		addDateButton = strings.TrimSpace(`<button name="add" hx-get="/date" hx-swap="outerHTML">Add Date</button>`)
	}

	span := createSpan(index, date)
	counter := createCounter(index, date)
	deleteButton := createDeleteButton(index)

	return fmt.Sprintf(entryHTML, span, counter, deleteButton, addDateButton)
}

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

	log.Printf("Name of triggering element (add/edit): %s\n", hxTriggerName)

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

	finalHTML, counterID, err := computeDateFinalHTML(date)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Initialize this counter element's map entry to 0.
	juicecount.Set(counterID, 0)
	log.Println(juicecount.Info())

	_, err = w.Write([]byte(finalHTML))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// computeDateFinalHTML returns the HTML to serve upon setting a date,
// dependent on whether the action (hxTriggerName) is an edit or an
// add. It also returns the ID of the new counter element, which is
// here used to initialize the corresponding map value to 0.
func computeDateFinalHTML(date string) (string, string, error) {
	if hxTriggerName != "edit" && hxTriggerName != "add" {
		return "", "", fmt.Errorf("Invalid HX trigger name: %s", hxTriggerName)
	}

	countElementIndex++
	newID := fmt.Sprintf("%s-%d", date, countElementIndex)
	wasAdd := (hxTriggerName == "add")
	return createEntry(countElementIndex, date, wasAdd), newID, nil
}
