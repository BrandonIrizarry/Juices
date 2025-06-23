package main

import (
	"html/template"
	"log"
	"os"
)

// FIXME: maybe there's a way we can combine this, dateInfo, and
// counts into an internal package that can then use tests?
type itemReport struct {
	itemName string
	date     string
}

// generateReports accumulates counts per category, per item, per
// date.
func generateReports() map[string]map[itemReport]int {
	reports := make(map[string]map[itemReport]int)

	// The ID (the map key) is ignored, since here is where we
	// group together counts belonging to the same date.
	for _, entry := range counts {
		itemName := entry.itemName
		date := entry.date
		count := entry.count

		ir := itemReport{itemName, date}

		category := categories[itemName]

		if reports[category] == nil {
			reports[category] = make(map[itemReport]int)
		}

		// If the item-with-date combination exists in the
		// inner map, add to the existing count there; else,
		// start a new entry.
		_, ok := reports[category][ir]

		if ok {
			reports[category][ir] += count
		} else {
			reports[category][ir] = count
		}
	}

	log.Printf("Accumulated reports: %v\n", reports)
	return reports
}

type dateInfo struct {
	Date  string
	Count int
}

// convertToHeadings creates the final view model used by the
// report.html template, where each item is grouped with the
// date-with-count combinations associated with it.
func convertToHeadings(reports map[string]map[itemReport]int) map[string]map[string][]dateInfo {
	headings := make(map[string]map[string][]dateInfo)

	for category, inner := range reports {
		log.Printf("Category: %v; inner: %v\n", category, inner)

		headings[category] = make(map[string][]dateInfo)

		for ir, count := range inner {
			itemName := ir.itemName

			if _, ok := headings[category][itemName]; !ok {
				headings[category][itemName] = make([]dateInfo, 0)
			}

			di := dateInfo{ir.date, count}
			headings[category][itemName] = append(headings[category][itemName], di)
		}
	}

	log.Printf("Generated headings: %v\n", headings)
	return headings
}

// writeReportsFile writes the template to disk using the appropriate
// view model.
//
// FIXME: can we combine this and 'convertToHeadings' into a single
// function, since all this does is a simple template fill-in?
func writeReportsFile(reportTemplate *template.Template, headings map[string]map[string][]dateInfo) error {
	reportFile, err := os.OpenFile("app/report.html", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)

	if err != nil {
		return err
	}

	defer reportFile.Close()

	if err := reportTemplate.Execute(reportFile, headings); err != nil {
		return err
	}

	return nil
}
