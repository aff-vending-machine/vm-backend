package request

type Create struct {
	BranchID     uint   `json:"branch_id" validate:"required"`
	Name         string `json:"name" validate:"required"`
	SerialNumber string `json:"serial_number" validate:"required"`
	Location     string `json:"location,omitempty"`
	Type         string `json:"type,omitempty"`
	Vendor       string `json:"vendor,omitempty"`
}
