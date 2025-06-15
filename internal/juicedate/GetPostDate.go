package juicedate

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/BrandonIrizarry/juices/internal/cid"
	"github.com/BrandonIrizarry/juices/internal/juicecount"
	"github.com/BrandonIrizarry/juices/internal/juicehtml"
)

// GetDate serves the HTML5 date widget to the page.
func GetDate(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving %s %s\n", r.Method, r.URL.Path)

	w.Header().Set("Content-Type", "text/html")

	// Whenever possible, we only work with IDs. In this case,
	// there is only one Add Item button per h1 header, and the
	// ids are differentiated precisely by the value of this
	// header.
	addButtonID := r.Header.Get("Hx-Trigger")

	log.Printf("ID of triggering element (add/edit): %s\n", addButtonID)

	// Note: since this could be an Add Date button, the only
	// valid fields for 'canonicalID' are 'WidgetType' and
	// 'ItemName'.
	canonicalID, err := cid.ParseCanonicalID(addButtonID)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	wtype := canonicalID.WidgetType

	if wtype != "add" && wtype != "edit" {
		message := fmt.Sprintf("Unexpected trigger element: %s", wtype)
		log.Println(message)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}

	// If an edit, delete the existing map entry.
	if wtype == "edit" {
		juicecount.Delete(canonicalID)
		log.Println(juicecount.Info())
	}

	itemName := canonicalID.ItemName
	getDateHTML := juicehtml.CreateGetDateHTML(wtype, itemName)

	_, err = w.Write([]byte(getDateHTML))

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

	canonicalID, err := cid.ParseCanonicalID(r.Header.Get("Hx-Trigger"))

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
