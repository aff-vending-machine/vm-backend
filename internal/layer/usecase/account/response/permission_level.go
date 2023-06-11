package response

type AccountPermission struct {
	Level    int  `json:"level"`
	BranchID uint `json:"branch_id"`
}
