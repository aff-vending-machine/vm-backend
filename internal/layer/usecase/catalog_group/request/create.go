package request

type Create struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	IsEnable    bool   `json:"is_enable" validate:"required"`
}
