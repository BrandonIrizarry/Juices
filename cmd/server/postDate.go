package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/BrandonIrizarry/juices/internal/cid"
	"github.com/BrandonIrizarry/juices/internal/juicecount"
	"github.com/BrandonIrizarry/juices/internal/juicehtml"
)

// PostDate serves a row consisting of the selected date, an HTML5
// counter widget, and a Delete button. If this is an Add Date
// operation, another Add Date button is appended to the served HTML.
func postDate(w http.ResponseWriter, r *http.Request) {
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

	canonicalID, err := cid.ParseCanonicalID(r)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	itemName := canonicalID.ItemName
	widgetType := canonicalID.WidgetType
	finalHTML, index, err := juicehtml.ComputeDateFinalHTML(date, itemName, widgetType)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Initialize this counter element's map entry to 0.
	juicecount.Set("count", itemName, date, index, 0)

	log.Println(juicecount.Info())

	_, err = w.Write([]byte(finalHTML))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
