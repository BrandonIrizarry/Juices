package main

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
