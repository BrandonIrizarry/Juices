package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BrandonIrizarry/juices/internal/cid"
	"github.com/BrandonIrizarry/juices/internal/juicecount"
	"github.com/BrandonIrizarry/juices/internal/juicehtml"
)

// GetDate serves the HTML5 date widget to the page.
func getDate(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving %s %s\n", r.Method, r.URL.Path)

	w.Header().Set("Content-Type", "text/html")

	// Note: since this could be an Add Date button, the only
	// valid fields for 'canonicalID' are 'WidgetType' and
	// 'ItemName'.
	canonicalID, err := cid.ParseCanonicalID(r)

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
		if err := juicecount.Delete(canonicalID); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Println(juicecount.Info())
	}

	itemName := canonicalID.ItemName
	getDateHTML := juicehtml.CreateGetDateHTML(wtype, itemName)

	_, err = w.Write([]byte(getDateHTML))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
