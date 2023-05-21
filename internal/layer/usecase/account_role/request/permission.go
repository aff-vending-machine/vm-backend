package request

type Permission struct {
	Scope string `json:"scope"`
	Level int    `json:"level"`
}
