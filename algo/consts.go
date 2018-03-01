package algo

// State handles.
const (
	Submitted = iota
	Acknowledged
	PathfinderInit
	PathfinderAsk
	PathfinderAcknowledge
	PathfinderDetermine
	SentToFullfillmentCenter
	FulfillmentCenterAcknowledge
	Fulfilling
	Shipped
)

const (
	// How much of the inventory is available.
	Inventory = iota
	// How fast the customer needs it
	Urgency
	// How expensive the last mile provider costs
	LastMile
	// How far away the FC is from the customer.
	Distance
)
