package main

import (
	"fmt"
	"os"
	"strings"

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
	date  string
	count int
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
	f, err := os.Create("app/reports.txt")

	if err != nil {
		return err
	}

	defer f.Close()

	for itemName, dis := range headings {
		realItemName := kebab.UndoKebabCase(itemName)

		if _, err := f.WriteString(realItemName + "\n" + strings.Repeat("-", len(realItemName)) + "\n"); err != nil {
			return err
		}

		for _, di := range dis {
			line := fmt.Sprintf("%d (%s)\n", di.count, di.date)

			if _, err := f.WriteString(line); err != nil {
				return err
			}
		}

		if _, err := f.WriteString("\n\n"); err != nil {
			return err
		}
	}

	return nil
}
