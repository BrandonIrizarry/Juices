package main

import (
	"log"
	"net/http"
)

func getReport(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving %s %s\n", r.Method, r.URL.Path)

	reports := generateReports()

	log.Printf("Reports: %v\n", reports)

	// Clear the map, since redirecting will erase all progress.
	clear(counts)

	http.Redirect(w, r, "/app", http.StatusSeeOther)
}
