package request

type Create struct {
	Name        string       `json:"name" validate:"required"`
	Permissions []Permission `json:"permissions" validate:"required"`
}
