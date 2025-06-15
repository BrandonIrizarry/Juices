package juicedate

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/BrandonIrizarry/juices/internal/juicecount"
	"github.com/BrandonIrizarry/juices/internal/juicehtml"
)

// The name of the widget requesting GET /date (could be "add" or "edit".)
var hxTriggerName string

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

	// If an edit, delete the existing map entry.
	if hxTriggerName == "edit" {
		id := r.Header.Get("Hx-Trigger")

		rawIndex, found := strings.CutPrefix(id, "edit-")

		if !found {
			message := fmt.Sprintf("Invalid hxTriggerName: %s", hxTriggerName)
			log.Println(message)
			http.Error(w, message, http.StatusInternalServerError)
			return
		}

		index, err := strconv.Atoi(rawIndex)

		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		juicecount.Delete(index)
	}

	getDateHTML := strings.TrimSpace(`<input type="date" name="date" hx-post="/date" hx-swap="outerHTML"/>`)
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

	finalHTML, counterID, err := juicehtml.ComputeDateFinalHTML(date, hxTriggerName)

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
