package hybrid_retail_store

// Defines an order from the customer.
type Order struct {
	// What the customer is ordering.
	Item               string
	// How many of the items is the customer ordering.
	Count              int
	// Where the item is being shipped.
	DestinationZipCode int
	// How expensive this item is to ship.
	Weight int
	// How soon they need it.
	Urgency int

	targetFulfillmentCenter FulFillmentCenter
}

