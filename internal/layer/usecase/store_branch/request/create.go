package request

type Create struct {
	Name     string `json:"name" validate:"required"`
	Location string `json:"location" validate:"required"`
	IsEnable bool   `json:"is_enable" validate:"required"`
}
