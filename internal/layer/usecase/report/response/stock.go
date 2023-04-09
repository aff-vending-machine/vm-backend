package response

import "time"

type Stock struct {
	Code     string    `json:"code"`
	Stock    int       `json:"stock"`
	Capacity int       `json:"capacity"`
	Sold     int       `json:"sold"`
	From     time.Time `json:"from"`
	To       time.Time `json:"to"`
	Price    float64   `json:"price"`
}

type CartItem struct {
	Code     string  `json:"code"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}
