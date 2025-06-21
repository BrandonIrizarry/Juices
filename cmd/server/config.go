package main

import (
	"errors"
	"fmt"
	"html/template"

	"github.com/BrandonIrizarry/juices/internal/kebab"
)

type config struct {
	views         map[string]*template.Template
	generateEntry func() (*template.Template, error)
}

func (cfg *config) initViews() error {
	if cfg.views != nil {
		return fmt.Errorf("views map already exists: %v", cfg.views)
	}

	// Define the main view (dashboard).
	cfg.views = make(map[string]*template.Template)

	fnMap := template.FuncMap{
		"kebabCase":     kebab.KebabCase,
		"undoKebabCase": kebab.UndoKebabCase,
		"createAcc":     createAcc,
		"inc":           createInc(),
	}

	start, err := template.New("start").Funcs(fnMap).ParseFiles("assets/start.html")

	if err != nil {
		return err
	}

	cfg.views["start"] = start

	// Prepare the report template.
	report, err := template.Must(start.Clone()).Funcs(fnMap).ParseFiles("assets/views/report.html")

	if err != nil {
		return err
	}

	cfg.views["report"] = report

	return nil
}

func (cfg *config) initEntryWithIndex() error {
	if cfg.generateEntry != nil {
		return errors.New("generateEntry is already initialized")
	}

	var index int

	cfg.generateEntry = func() (*template.Template, error) {
		entryHTML, err := template.New("entry").Funcs(template.FuncMap{
			"inc": func() int {
				index++
				return index
			},
		}).ParseFiles("assets/entry.html")

		if err != nil {
			return nil, err
		}

		return entryHTML, nil
	}

	return nil
}

type countInfo struct {
	Count int
	Total int
}

// FIXME: can we handle the closure update from this file, instead of
// from within the template (as we do with "inc")?
func createAcc() func(count int) countInfo {
	info := countInfo{
		Count: 0,
		Total: 0,
	}

	return func(count int) countInfo {
		info = countInfo{count, info.Total + count}

		return info
	}
}

func createInc() func() int {
	var index int

	return func() int {
		index++
		return index
	}
}
