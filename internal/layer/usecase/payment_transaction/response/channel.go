package response

type Channel struct {
	Name     string `json:"name"`
	Channel  string `json:"channel"`
	Vendor   string `json:"vendor"`
	IsEnable bool   `json:"is_enable"`
}
