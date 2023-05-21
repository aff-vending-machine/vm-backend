package response

type Slot struct {
	ID        uint     `json:"id"`
	ProductID uint     `json:"product_id"`
	Product   *Product `json:"product,omitempty"`
	Code      string   `json:"code"`
	Stock     int      `json:"stock"`
	Capacity  int      `json:"capacity"`
	IsEnable  bool     `json:"is_enable"`
}
