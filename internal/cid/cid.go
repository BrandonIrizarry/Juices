package cid

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type CanonicalID struct {
	ItemName   string
	WidgetType string
	Date       string
	Index      int
}

func ParseCanonicalID(r *http.Request) (CanonicalID, error) {
	rawCID := r.Header.Get("Hx-Trigger")

	if rawCID == "" {
		return CanonicalID{}, errors.New("Empty Hx-Trigger header (missing id attribute)")
	}

	parts := strings.Split(rawCID, "_")

	if len(parts) != 4 {
		return CanonicalID{}, fmt.Errorf("Malformed CID: %s", rawCID)
	}

	widgetType := parts[0]
	itemName := parts[1]
	date := parts[2]
	rawIndex := parts[3]

	index, err := strconv.Atoi(rawIndex)

	if err != nil {
		return CanonicalID{}, err
	}

	CID := CanonicalID{itemName, widgetType, date, index}

	return CID, nil
}
