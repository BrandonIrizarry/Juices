package juicecount

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var idsToCounts = make(map[string]int)

func Set(id string, count int) error {
	if id == "" {
		return errors.New("Empty ID")
	}

	if count < 0 {
		return fmt.Errorf("Invalid count: %d", count)
	}

	idsToCounts[id] = count

	return nil
}

func Get(id string) (int, error) {
	count, ok := idsToCounts[id]

	if !ok {
		return 0, fmt.Errorf("Nonexistent ID: %s", id)
	}

	return count, nil
}

func Delete(idIndex int) error {
	var foundID string

	for id := range idsToCounts {
		parts := strings.SplitN(id, "-", 3)

		if len(parts) != 3 {
			return fmt.Errorf("Invalid map-key/counter ID: %s", id)
		}

		currentIDIndex, err := strconv.Atoi(parts[2])

		if err != nil {
			return err
		}

		if currentIDIndex == idIndex {
			foundID = id
			break
		}
	}

	delete(idsToCounts, foundID)

	return nil
}

func Info() string {
	return fmt.Sprintf("Current counts: %v\n", idsToCounts)
}
