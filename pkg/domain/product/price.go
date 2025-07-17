package product

type Price struct {
	Original           int     `json:"original"`
	Final              int     `json:"final"`
}

type PriceResponse struct {
	Original           int     `json:"original"`
	Final              int     `json:"final"`
	DiscountPercentage *string `json:"discount_percentage"`
	Currency           string  `json:"currency"`
}
