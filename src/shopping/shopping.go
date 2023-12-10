package shopping

import (
	"little-go/src/shopping/db"
)

func PriceCheck11(id int) (float64, bool) {
	item := db.LoadItem(id)
	if item == nil {
		return 0.0, false
	}
	return item.Price, true
}

func PriceCheck(id int) (float64, bool) {
	item := db.LoadItem(id)
	if item == nil {
		return 0.0, false
	}
	return item.Price, true
}

func PriceCheck2(id int) (float64, bool) {
	item := db.LoadItem(id)
	if item == nil {
		return 0.0, false
	}
	return item.Price, true
}
