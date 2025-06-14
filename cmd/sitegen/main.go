package main

import (
	"bufio"
	"html/template"
	"log"
	"os"
	"strings"
)

func main() {
	items, err := inventoryItems()

	if err != nil {
		log.Fatal(err)
	}

	t, err := template.ParseFiles("assets/template.html")

	if err != nil {
		log.Fatal(err)
	}

	t.Execute(os.Stdout, items)
}

func inventoryItems() ([]string, error) {
	file, err := os.Open("assets/inventory.txt")

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	buffer := make([]string, 0)

	for scanner.Scan() {
		item := strings.TrimSpace(scanner.Text())

		if item != "" {
			buffer = append(buffer, item)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return buffer, nil
}
