package account_token

import "vm-backend/configs"

type managementImpl struct {
	secretAccess    string
	secretRefresh   string
	audience        string
	issuer          string
	authorizedParty string
}

func NewManagement(cfg configs.JWTConfig) *managementImpl {
	return &managementImpl{
		secretAccess:    cfg.Reference1,
		secretRefresh:   cfg.Reference2,
		audience:        cfg.Audience,
		issuer:          cfg.Issuer,
		authorizedParty: cfg.AuthorizedParty,
	}
}
