package juicecount

import (
	"fmt"

	"github.com/BrandonIrizarry/juices/internal/cid"
)

var idsToCounts = make(map[cid.CanonicalID]int)

func Set(widgetType, itemName, date string, index, count int) {
	cid := cid.CanonicalID{
		WidgetType: widgetType,
		ItemName:   itemName,
		Date:       date,
		Index:      index,
	}

	idsToCounts[cid] = count
}

func SetBulk(cid cid.CanonicalID, count int) {
	idsToCounts[cid] = count
}

func Get(cid cid.CanonicalID) (int, error) {
	count, ok := idsToCounts[cid]

	if !ok {
		return 0, fmt.Errorf("Nonexistent ID: %#v", cid)
	}

	return count, nil
}

func Delete(cid cid.CanonicalID) {
	delete(idsToCounts, cid)
}

func Info() string {
	return fmt.Sprintf("Current counts: %#v\n", idsToCounts)
}
