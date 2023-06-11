package account

import (
	"context"
	"time"
	"vm-backend/internal/core/domain/account"
	"vm-backend/internal/layer/usecase/account/response"
	"vm-backend/pkg/helpers/db"
)

type usecase struct {
	roleRepo    account.RoleRepository
	userRepo    account.UserRepository
	passwordMgr account.PasswordManagement
	tokenMgr    account.TokenManagement
}

func NewUsecase(
	arr account.RoleRepository,
	aur account.UserRepository,
	apm account.PasswordManagement,
	atm account.TokenManagement,
) account.Usecase {
	return &usecase{
		arr,
		aur,
		apm,
		atm,
	}
}

func (uc *usecase) generateToken(ctx context.Context, user *account.User) (*response.AuthResult, error) {
	at := account.NewAccessToken(*user)
	accessToken, err := uc.tokenMgr.CreateAccessToken(ctx, at)
	if err != nil {
		return nil, err
	}

	rt := account.NewRefreshToken(*user)
	refreshToken, err := uc.tokenMgr.CreateRefreshToken(ctx, rt)
	if err != nil {
		return nil, err
	}

	query := db.NewQuery().AddWhere("id = ?", user.ID)
	data := map[string]interface{}{
		"last_login": time.Now(),
		"last_token": accessToken,
	}
	_, err = uc.userRepo.Update(ctx, query, data)
	if err != nil {
		return nil, err
	}

	return &response.AuthResult{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User: response.User{
			ID:       user.ID,
			Username: user.Username,
			Role:     user.Role.Name,
		},
	}, nil
}
