package hybrid_retail_store

const (
	FC1Location = 80021
	FC2Location = 98121
)

// Defines our regional fulfillment center.
type FulFillmentCenter struct {
	CurrentOrder Order
	Destination int
}
