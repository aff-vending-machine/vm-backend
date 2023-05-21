package models

type Channel struct {
	Channel    string `json:"channel"`     // primary key
	MerchantID uint   `json:"merchant_id"` // primary key
	Name       string `json:"name"`
	Vendor     string `json:"vendor"`
	IsEnable   bool   `json:"is_enable"`
	Host       string `json:"host"`
	BillerCode string `json:"biller_code"`
	BillerID   string `json:"biller_id"`
	Token      string `json:"token"`
	StoreID    string `json:"store_id"`
	TerminalID string `json:"terminal_id"`
}

func (m *Channel) ToUpdate() map[string]interface{} {
	return map[string]interface{}{
		"name":        m.Name,
		"vendor":      m.Vendor,
		"is_enable":   m.IsEnable,
		"host":        m.Host,
		"biller_code": m.BillerCode,
		"biller_id":   m.BillerID,
		"token":       m.Token,
		"store_id":    m.StoreID,
		"terminal_id": m.TerminalID,
	}
}
