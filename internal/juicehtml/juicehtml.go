package juicehtml

import (
	"fmt"
	"strings"
)

// The index appended to the computed ID of a counter widget (input
// type=number.) This is used to enforce unique HTML id attribute
// values among the counter widgets themselves.
var countElementIndex = 0

// computeDateFinalHTML returns the HTML to serve upon setting a date,
// dependent on whether the action (hxTriggerName) is an edit or an
// add. It also returns the ID of the new counter element, which is
// here used to initialize the corresponding map value to 0.
func ComputeDateFinalHTML(date, hxTriggerName string) (string, string, error) {
	if hxTriggerName != "edit" && hxTriggerName != "add" {
		return "", "", fmt.Errorf("Invalid HX trigger name: %s", hxTriggerName)
	}

	countElementIndex++
	newID := fmt.Sprintf("%s-%d", date, countElementIndex)
	wasAdd := (hxTriggerName == "add")
	return createEntry(countElementIndex, date, wasAdd), newID, nil
}

func createSpan(index int, date string) string {
	spanHTML := strings.TrimSpace(`
<span name="edit" id="edit-%d" hx-get="/date" hx-swap="outerHTML" hx-target="closest div">%s</span>
`)

	return fmt.Sprintf(spanHTML, index, date)
}

func createCounter(index int, date string) string {
	counterHTML := strings.TrimSpace(`
<input type="number" name="count" hx-post="/count" hx-trigger="change delay:1s" min=0 id="%s-%d" value="0" />
`)
	return fmt.Sprintf(counterHTML, date, index)
}

func createDeleteButton(index int) string {
	deleteButtonHTML := strings.TrimSpace(`
<button id="delete-%d" hx-delete="/date" hx-swap="delete" hx-target="closest div" hx-confirm="Delete this row?">Delete</button>
`)
	return fmt.Sprintf(deleteButtonHTML, index)
}

func createEntry(index int, date string, wasAdd bool) string {
	entryHTML := strings.TrimSpace(`
<div class="entry">
%s
%s
%s
</div>
%s`)

	var addDateButton string

	if wasAdd {
		addDateButton = strings.TrimSpace(`<button name="add" hx-get="/date" hx-swap="outerHTML">Add Date</button>`)
	}

	span := createSpan(index, date)
	counter := createCounter(index, date)
	deleteButton := createDeleteButton(index)

	return fmt.Sprintf(entryHTML, span, counter, deleteButton, addDateButton)
}
