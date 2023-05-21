package response

type Item struct {
	Code     string  `json:"code"`
	Name     string  `json:"name"`
	ImageUrl string  `json:"image_url"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	Received int     `json:"received"`
}
