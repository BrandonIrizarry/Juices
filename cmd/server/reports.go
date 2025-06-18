package main

import (
	"html/template"
	"os"

	"github.com/BrandonIrizarry/juices/internal/kebab"
)

type itemReport struct {
	itemName string
	date     string
}

func generateReports() map[itemReport]int {
	reports := make(map[itemReport]int)

	for _, entry := range counts {
		itemName := entry.itemName
		date := entry.date
		count := entry.count

		ir := itemReport{itemName, date}

		_, ok := reports[ir]

		if ok {
			reports[ir] += count
		} else {
			reports[ir] = count
		}
	}

	return reports
}

type dateInfo struct {
	Date  string
	Count int
}

func convertToHeadings(reports map[itemReport]int) map[string][]dateInfo {
	headings := make(map[string][]dateInfo)

	for ir, count := range reports {
		itemName := ir.itemName

		if _, ok := headings[itemName]; !ok {
			headings[itemName] = make([]dateInfo, 0)
		}

		di := dateInfo{ir.date, count}
		headings[itemName] = append(headings[itemName], di)
	}

	return headings
}

func writeReportsFile(headings map[string][]dateInfo) error {
	reportFile, err := os.OpenFile("app/report.html", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)

	if err != nil {
		return err
	}

	defer reportFile.Close()

	type countInfo struct {
		Count int
		Total int
	}

	createAcc := func() func(count int) countInfo {
		info := countInfo{
			Count: 0,
			Total: 0,
		}

		return func(count int) countInfo {
			info = countInfo{count, info.Total + count}

			return info
		}
	}

	// Prepare the report template.
	t, err := template.New("start").Funcs(template.FuncMap{
		// kebab.KebabCase is included here because we're
		// reusing the start template, which contains a
		// template block whose default definition uses this
		// function.
		"kebabCase":     kebab.KebabCase,
		"undoKebabCase": kebab.UndoKebabCase,
		"createAcc":     createAcc,
	}).ParseFiles("assets/start.html", "assets/report.html")

	if err != nil {
		return err
	}

	if err := t.ExecuteTemplate(reportFile, "start", headings); err != nil {
		return err
	}

	return nil
}
