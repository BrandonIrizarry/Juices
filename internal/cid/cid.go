package cid

import (
	"fmt"
	"strconv"
	"strings"
)

type CanonicalID struct {
	ItemName   string
	WidgetType string
	Date       string
	Index      int
}

func ParseCanonicalID(rawCID string) (CanonicalID, error) {
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
