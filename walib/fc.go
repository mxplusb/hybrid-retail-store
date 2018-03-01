package hybrid_retail_store

import "time"

const (
	FC1Location = 80021
	FC2Location = 98121
)

// Defines our regional fulfillment center.
type FulFillmentCenter struct {
	Destination int
}

// Defines an item in inventory.
type InventorySchema struct {
	ItemAdded     time.Time
	Item          string
	Inventory     int
	PurchaseCount int64
	Price         float64
}
