package main

import (
	"log"
	"net/http"
)

func (cfg *config) getReport(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving %s %s\n", r.Method, r.URL.Path)

	reports := generateReports()

	log.Printf("Reports: %v\n", reports)

	headings := convertToHeadings(reports)

	log.Printf("Headings: %v\n", headings)

	if err := writeReportsFile(cfg.views["report"], headings); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/app/report.html", http.StatusSeeOther)
}
