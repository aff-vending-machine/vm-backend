package response

type Branch struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	IsEnable bool   `json:"is_enable"`
}
