package preload

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/user"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/user/request"
	"github.com/aff-vending-machine/vm-backend/pkg/boot"
	"github.com/aff-vending-machine/vm-backend/pkg/errs"
	"github.com/rs/zerolog/log"
)

func CreateSuperAdmin(uc user.Usecase, roleID uint) uint {
	ctx := context.TODO()

	username := "superadmin"

	filter := &request.Filter{Username: &username}
	list, err := uc.List(ctx, filter)
	if errs.Not(err, "not found") {
		boot.TerminateWhenError(err, filter)
	}
	if err == nil && len(list) > 0 {
		return list[0].ID
	}

	log.Info().Msg("create super admin user")
	request := &request.Create{
		Username:  username,
		Password:  username,
		RoleID:    roleID,
		CreatedBy: "system",
	}
	id, err := uc.Create(ctx, request)
	boot.TerminateWhenError(err, request)

	return id
}

func CreateAdmin(uc user.Usecase, roleID uint) uint {
	ctx := context.TODO()

	username := "admin"

	filter := &request.Filter{Username: &username}
	list, err := uc.List(ctx, filter)
	if errs.Not(err, "not found") {
		boot.TerminateWhenError(err, filter)
	}
	if err == nil && len(list) > 0 {
		return list[0].ID
	}

	log.Info().Msg("create admin user")
	request := &request.Create{
		Username:  username,
		Password:  username,
		RoleID:    roleID,
		CreatedBy: "system",
	}
	id, err := uc.Create(ctx, request)
	boot.TerminateWhenError(err, request)

	return id
}
