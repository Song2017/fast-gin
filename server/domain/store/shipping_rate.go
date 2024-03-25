package order_domain

type ShippingRate struct {
	Type     string  `json:"type"`
	Currency string  `json:"currency"`
	Value    float32 `json:"value"`
	// WeightLimits string `json:"weightLimits"`
	// MinimumSpend string `json:"minimumSpend"`
}
