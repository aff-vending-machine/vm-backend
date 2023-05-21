package response

type Permission struct {
	Scope string `json:"scope"`
	Level int    `json:"level"`
}
