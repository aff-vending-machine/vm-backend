package auth_usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/aff-vending-machine/vm-backend/internal/core/domain/entity"
	"github.com/aff-vending-machine/vm-backend/internal/core/domain/jwt"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/auth/response"
)

func (uc *usecase) generateToken(ctx context.Context, user *entity.User) (*response.AuthResult, error) {
	at := jwt.NewAccessToken(*user)
	accessToken, err := uc.tokenMgr.CreateAccessToken(ctx, at)
	if err != nil {
		return nil, err
	}

	rt := jwt.NewRefreshToken(*user)
	refreshToken, err := uc.tokenMgr.CreateRefreshToken(ctx, rt)
	if err != nil {
		return nil, err
	}

	filter := []string{fmt.Sprintf("id||=||%d", user.ID)}
	data := map[string]interface{}{"last_login": time.Now()}
	_, err = uc.userRepo.UpdateMany(ctx, filter, data)
	if err != nil {
		return nil, err
	}

	return &response.AuthResult{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         response.ToUser(user),
	}, nil
}
