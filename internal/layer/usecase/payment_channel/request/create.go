package request

type Create struct {
	Name         string `json:"name" validate:"required"`
	Channel      string `json:"channel" validate:"required"`
	Vendor       string `json:"vendor"`
	IsEnable     bool   `json:"is_enable"`
	Host         string `json:"host"`
	MerchantID   string `json:"merchant_id"`
	MerchantName string `json:"merchant_name"`
	BillerCode   string `json:"biller_code"`
	BillerID     string `json:"biller_id"`
	StoreID      string `json:"store_id"`
	TerminalID   string `json:"terminal_id"`
}
