package response

type Stock struct {
	Code       string  `json:"code"`
	Name       string  `json:"name"`
	Sold       int     `json:"sold"`
	Price      float64 `json:"price"`
	TotalPrice float64 `json:"total_price"`
}

type CartItem struct {
	Code     string  `json:"code"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}
