package main

import (
	"fmt"
	"strings"
)

func nonEmptyValue(value string) (string, error) {
	if value == "" {
		return "", fmt.Errorf("Missing '%s' from path values", value)
	}

	return value, nil
}

func parseDate(dateFormValue string) (string, error) {
	parts := strings.SplitN(dateFormValue, "-", 2)

	if len(parts) != 2 {
		return "", fmt.Errorf("Invalid date: %s", dateFormValue)
	}

	return parts[1], nil
}
