package auth_usecase

import (
	"context"
	"fmt"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/auth/request"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/auth/response"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecase) GetPermissionLevel(ctx context.Context, req *request.GetPermissionLevel) (*response.PermissionLevel, error) {
	if v := validate.Struct(req); !v.Validate() {
		return nil, errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	user, err := uc.userRepo.FindOne(ctx, req.ToUserFilter())
	if err != nil {
		return nil, errors.Wrap(err, "unable to find one user")
	}
	if user.HasRole("") {
		return nil, fmt.Errorf("user has no rule")
	}

	role, err := uc.roleRepo.FindOne(ctx, req.ToRoleFilter(user.RoleID))
	if err != nil {
		return nil, errors.Wrap(err, "unable to find one role")
	}

	permission := role.HasPermission(req.Scope)
	if permission == 0 {
		return nil, fmt.Errorf("no permission to access")
	}

	return &response.PermissionLevel{Level: permission}, nil
}
