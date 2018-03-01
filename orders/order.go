package hybrid_retail_store

// Defines an order from the customer.
type Order struct {
	// What the customer is ordering.
	Item               string `json:"item"`
	// How many of the items is the customer ordering.
	Count              int `json:"count"`
	// Where the item is being shipped.
	DestinationZipCode int `json:"destination_zip_code"`
	// How expensive this item is to ship.
	Weight int `json:"weight"`
	// How soon they need it.
	Urgency int `json:"urgency"`
}
