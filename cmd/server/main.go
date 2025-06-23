package main

import (
	"bufio"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Set up server.
	mux := http.NewServeMux()

	var cfg config

	if err := cfg.initViews(); err != nil {
		log.Fatal(err)
	}

	if err := cfg.initEntryWithIndex(); err != nil {
		log.Fatal(err)
	}

	items, err := inventoryItems()

	if err != nil {
		log.Fatal(err)
	}

	if err := initIndexHTML(cfg.views["start"], items); err != nil {
		log.Fatal(err)
	}

	// Serve page as a static asset (includes CSS and JS, where
	// HTMX resides)
	serveMainPage := http.StripPrefix("/app/", http.FileServer(http.Dir("./app")))
	mux.Handle("GET /app/", serveMainPage)
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/app", http.StatusSeeOther)
	})

	// Endpoints (each handler is defined in its own file inside
	// the main package.
	mux.HandleFunc("GET /date/{itemName}", getDate)
	mux.HandleFunc("POST /date/{itemName}", cfg.postDate)
	mux.HandleFunc("DELETE /date", deleteDate)
	mux.HandleFunc("POST /count/{itemName}/{date}", postCount)
	mux.HandleFunc("GET /report", cfg.getReport)
	mux.HandleFunc("GET /prepare", getPrepare)
	mux.HandleFunc("POST /clear", postClear)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Fatal(srv.ListenAndServe())
}

// Map an item name to the category it belongs to.
var categories = make(map[string]string)
var registeredCategories = make(map[string][]string)
var categoryOrder = make(map[string]int)

func inventoryItems() ([]string, error) {
	file, err := os.Open("assets/inventory.txt")

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// The return value, used to fill in the start.html template.
	buffer := make([]string, 0)

	var category string
	var categoryIndex int

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, "*") {
			// The line is a category.
			category = strings.TrimSpace(line[1:])

			if registeredCategories[category] == nil {
				registeredCategories[category] = make([]string, 0)
			}

			categoryIndex++
			categoryOrder[category] = categoryIndex
		} else if line != "" {
			// The line is an item; use a new variable for
			// readability.
			item := line

			if category == "" {
				return nil, errors.New("Current category is unset")
			} else if registeredCategories[category] == nil {
				return nil, fmt.Errorf("Undefined category: %s", category)
			}

			categories[item] = category
			registeredCategories[category] = append(registeredCategories[category], item)

			buffer = append(buffer, item)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	log.Printf("Items under categories: %v\n", categories)
	log.Printf("Registered categories: %v\n", registeredCategories)
	log.Printf("Order of categories: %v\n", categoryOrder)
	return buffer, nil
}

// initIndexHTML copies the start template to 'app/index.html', so
// that the file server can pick it up.
func initIndexHTML(start *template.Template, items []string) error {
	indexHTML, err := os.OpenFile("app/index.html", os.O_RDWR, 0644)

	if err != nil {
		return err
	}

	defer indexHTML.Close()

	if err := start.Execute(indexHTML, items); err != nil {
		return err
	}

	return nil
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
