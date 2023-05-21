package response

type Machine struct {
	Name          string `json:"name"`
	SerialNumber  string `json:"serial_number"`
	Location      string `json:"location"`
	Type          string `json:"type"`
	Vendor        string `json:"vendor"`
	Status        string `json:"status"`
	RegisterCount int    `json:"register_count"`
}
