package main

import (
	"fmt"
	"os"
	"strings"
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

func writeReportsFile(reports map[itemReport]int) error {
	f, err := os.Create("app/reports.txt")

	if err != nil {
		return err
	}

	defer f.Close()

	for ir, count := range reports {
		realItemName := undoKebabCase(ir.itemName)
		date := ir.date

		line := fmt.Sprintf("%s: %d (%s)\n", realItemName, count, date)

		if _, err := f.WriteString(line); err != nil {
			return err
		}
	}

	return nil
}

func undoKebabCase(itemName string) string {
	parts := strings.Split(itemName, "-")

	for i := range parts {
		parts[i] = strings.ToUpper(parts[i])
	}

	return strings.Join(parts, " ")
}
