package response

type Healthy struct {
	Ready   bool   `json:"ready"`
	Message string `json:"message"`
}
