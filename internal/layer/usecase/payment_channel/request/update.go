package request

import (
	"vm-backend/pkg/conv"
	"vm-backend/pkg/db"
)

type Update struct {
	ID           uint    `json:"id" query:"id" validate:"required"`
	Name         *string `json:"name,omitempty"`
	Channel      *string `json:"channel,omitempty"`
	Vendor       *string `json:"vendor,omitempty"`
	IsEnable     *bool   `json:"is_enable,omitempty"`
	Host         *string `json:"host,omitempty"`
	MerchantID   *string `json:"merchant_id,omitempty"`
	MerchantName *string `json:"merchant_name,omitempty"`
	BillerCode   *string `json:"biller_code,omitempty"`
	BillerID     *string `json:"biller_id,omitempty"`
	Token        *string `json:"token,omitempty"`
	StoreID      *string `json:"store_id,omitempty"`
	TerminalID   *string `json:"terminal_id,omitempty"`
}

func (r *Update) ToQuery() *db.Query {
	return db.NewQuery().
		AddWhere("id = ?", r.ID)
}

func (r *Update) ToUpdate() map[string]interface{} {
	result, _ := conv.StructToMap(r)
	delete(result, "id")
	return result
}
