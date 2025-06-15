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

	t, err := template.New("template").Funcs(template.FuncMap{
		"kebabCase": func(name string) string {
			subwords := strings.Fields(name)

			for i := range subwords {
				subwords[i] = strings.ToLower(subwords[i])
			}

			return strings.Join(subwords, "-")
		},
	}).ParseFiles("assets/template.html")

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
