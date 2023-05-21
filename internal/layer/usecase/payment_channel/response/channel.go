package response

type Channel struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Channel      string `json:"channel"`
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
