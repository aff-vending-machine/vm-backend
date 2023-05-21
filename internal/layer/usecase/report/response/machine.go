package response

type Machine struct {
	ID            uint               `json:"id"`
	Name          string             `json:"name"`
	SerialNumber  string             `json:"serial_number"`
	TotalPayments map[string]float64 `json:"total_payments"`
	Total         float64            `json:"total"`
}
