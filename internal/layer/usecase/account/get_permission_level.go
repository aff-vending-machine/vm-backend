package account

import (
	"context"
	"fmt"
	"vm-backend/internal/layer/usecase/account/request"
	"vm-backend/internal/layer/usecase/account/response"

	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecase) GetPermissionLevel(ctx context.Context, req *request.GetPermissionLevel) (*response.PermissionLevel, error) {
	if v := validate.Struct(req); !v.Validate() {
		err := v.Errors.OneError()
		log.Error().Err(err).Interface("request", req).Msg("unable to validate request")
		return nil, errors.Wrap(err, "unable to validate request")
	}

	query := req.ToUserQuery()
	user, err := uc.userRepo.FindOne(ctx, query)
	if err != nil {
		log.Error().Err(err).Interface("query", query).Msg("unable to find one user")
		return nil, errors.Wrap(err, "unable to find one user")
	}

	if user.HasRole("") {
		log.Error().Interface("query", query).Interface("user", user).Msg("user has no role")
		return nil, errors.New("user has no role")
	}

	permission := user.Role.HasPermission(req.Scope)
	if permission == 0 {
		log.Error().Str("scope", req.Scope).Interface("user", user).Msg("user has no permission")
		return nil, fmt.Errorf("user has no permission")
	}

	return &response.PermissionLevel{Level: permission}, nil
}
