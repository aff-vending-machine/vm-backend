package account_password

import "vm-backend/configs"

type managementImpl struct {
	salt int
}

func NewManagement(cfg configs.BCryptConfig) *managementImpl {
	return &managementImpl{
		cfg.Salt,
	}
}
