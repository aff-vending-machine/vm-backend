package request

type Create struct {
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
	CreatedBy string `json:"created_by" validate:"required"`
	RoleID    uint   `json:"role_id" validate:"required"`
	BranchID  *uint  `json:"branch_id,omitempty"`
}
