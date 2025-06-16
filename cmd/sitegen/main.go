package main

import (
	"bufio"
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/BrandonIrizarry/juices/internal/kebab"
)

func main() {
	items, err := inventoryItems()

	if err != nil {
		log.Fatal(err)
	}

	fnMap := template.FuncMap{"kebabCase": kebab.KebabCase}
	t, err := template.New("start").Funcs(fnMap).ParseFiles("assets/start.html")

	if err != nil {
		log.Fatal(err)
	}

	if err := t.Execute(os.Stdout, items); err != nil {
		log.Fatal(err)
	}
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
