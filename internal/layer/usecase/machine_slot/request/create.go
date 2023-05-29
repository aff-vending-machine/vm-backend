package request

type Create struct {
	MachineID        uint   `json:"machine_id" query:"machine_id" validate:"required"`
	CatalogProductID uint   `json:"catalog_product_id,omitempty"`
	Code             string `json:"code" validate:"required"`
	Stock            int    `json:"stock,omitempty" validate:"min:0"`
	Capacity         int    `json:"capacity,omitempty" validate:"min:0"`
	IsEnable         bool   `json:"is_enable,omitempty"`
}
