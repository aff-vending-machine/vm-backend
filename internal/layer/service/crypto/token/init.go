package token

import "github.com/aff-vending-machine/vm-backend/config"

type managerImpl struct {
	SecretAccess    string
	SecretRefresh   string
	audience        string
	issuer          string
	authorizedParty string
}

func New(config config.JWTConfig) *managerImpl {
	return &managerImpl{
		SecretAccess:    config.Reference1,
		SecretRefresh:   config.Reference2,
		audience:        config.Audience,
		issuer:          config.Issuer,
		authorizedParty: config.AuthorizedParty,
	}
}
