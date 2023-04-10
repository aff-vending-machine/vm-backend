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

type SortStockByCode []Stock

func (p SortStockByCode) Len() int           { return len(p) }
func (p SortStockByCode) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p SortStockByCode) Less(i, j int) bool { return p[i].Code < p[j].Code }
