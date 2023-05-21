package response

type Group struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsEnable    bool   `json:"is_enable"`
}
