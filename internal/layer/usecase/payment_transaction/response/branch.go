package response

type Branch struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	IsEnable bool   `json:"is_enable"`
}
