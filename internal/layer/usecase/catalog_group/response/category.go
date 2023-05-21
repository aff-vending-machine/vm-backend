package response

type Group struct {
	ID          uint      `json:"id"`
	Products    []Product `json:"products"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsEnable    bool      `json:"is_enable"`
}
