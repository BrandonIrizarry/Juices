package juicehtml

import (
	"fmt"
	"strings"
)

// The index appended to the computed ID of a counter widget (input
// type=number.) This is used to enforce unique HTML id attribute
// values among the counter widgets themselves.
var countElementIndex = 0

func createSpan(index int, itemName, date string) string {
	CID := generateRawCanonicalID(index, "edit", itemName, date)

	spanHTML := strings.TrimSpace(`
<span name="edit" id="%s" hx-get="/date" hx-swap="outerHTML" hx-target="closest div">%s</span>
`)

	return fmt.Sprintf(spanHTML, CID, date)
}

func createCounter(index int, itemName, date string) string {
	CID := generateRawCanonicalID(index, "count", itemName, date)

	counterHTML := strings.TrimSpace(`
<input type="number" name="count" id="%s" hx-post="/count" hx-trigger="change delay:1s" min=0 value="0" />
`)
	return fmt.Sprintf(counterHTML, CID)
}

func createDeleteButton(index int, itemName, date string) string {
	CID := generateRawCanonicalID(index, "delete", itemName, date)

	deleteButtonHTML := strings.TrimSpace(`
<button id="%s" hx-delete="/date" hx-swap="delete" hx-target="closest div" hx-confirm="Delete this row?">Delete</button>
`)
	return fmt.Sprintf(deleteButtonHTML, CID)
}

func createAddDateButton(itemName string) string {
	addDateButton := strings.TrimSpace(`<button name="add" id="%s" hx-get="/date" hx-swap="outerHTML">Add Date</button>`)

	return fmt.Sprintf(addDateButton, itemName)
}

func createEntry(index int, itemName, date string, wasAdd bool) string {
	entryHTML := strings.TrimSpace(`
<div class="entry">
%s
%s
%s
</div>
%s`)

	var addDateButton string

	if wasAdd {
		addDateButton = createAddDateButton(itemName)
	}

	span := createSpan(index, itemName, date)
	counter := createCounter(index, itemName, date)
	deleteButton := createDeleteButton(index, itemName, date)

	return fmt.Sprintf(entryHTML, span, counter, deleteButton, addDateButton)
}

func generateRawCanonicalID(index int, widgetType, itemName, date string) string {
	return fmt.Sprintf("%s_%s_%s_%d", widgetType, itemName, date, index)
}

// computeDateFinalHTML returns the HTML to serve upon setting a date,
// dependent on whether the action (hxTriggerName) is an edit or an
// add. It also returns the ID of the new counter element, which is
// here used to initialize the corresponding map value to 0.
func ComputeDateFinalHTML(date, itemName, widgetType string) (string, int, error) {
	if widgetType != "edit" && widgetType != "add" {
		return "", 0, fmt.Errorf("Invalid HX trigger name: %s", widgetType)
	}

	countElementIndex++
	wasAdd := (widgetType == "add")
	return createEntry(countElementIndex, itemName, date, wasAdd), countElementIndex, nil
}
