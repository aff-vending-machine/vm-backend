package preload

import (
	"context"

	"github.com/aff-vending-machine/vm-backend/internal/core/domain/permission"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/role"
	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase/role/request"
	"github.com/aff-vending-machine/vm-backend/pkg/boot"
	"github.com/aff-vending-machine/vm-backend/pkg/errs"
	"github.com/rs/zerolog/log"
)

func CreateSuperAdminRole(uc role.Usecase) uint {
	ctx := context.TODO()

	role := "super-admin"

	filter := &request.Filter{Name: &role}
	list, err := uc.List(ctx, filter)
	if errs.Not(err, "not found") {
		boot.TerminateWhenError(err, filter)
	}
	if err == nil && len(list) > 0 {
		return list[0].ID
	}

	log.Info().Msg("create super admin role")
	request := &request.Create{
		Name: role,
		Permissions: []request.Permission{
			{Scope: "machine", Level: permission.Admin},
			{Scope: "machine_slot", Level: permission.Admin},
			{Scope: "payment_channel", Level: permission.Admin},
			{Scope: "product", Level: permission.Admin},
			{Scope: "report", Level: permission.Admin},
			{Scope: "role", Level: permission.Admin},
			{Scope: "sync", Level: permission.Admin},
			{Scope: "transaction", Level: permission.Admin},
			{Scope: "user", Level: permission.Admin},
		},
	}
	id, err := uc.Create(ctx, request)
	boot.TerminateWhenError(err, request)

	return id
}

func CreateAdminRole(uc role.Usecase) uint {
	ctx := context.TODO()

	role := "admin"

	filter := &request.Filter{Name: &role}
	list, err := uc.List(ctx, filter)
	if errs.Not(err, "not found") {
		boot.TerminateWhenError(err, filter)
	}
	if err == nil && len(list) > 0 {
		return list[0].ID
	}

	log.Info().Msg("create admin role")
	request := &request.Create{
		Name: role,
		Permissions: []request.Permission{
			{Scope: "machine", Level: permission.Editor},
			{Scope: "machine_slot", Level: permission.Editor},
			{Scope: "payment_channel", Level: permission.Editor},
			{Scope: "product", Level: permission.Editor},
			{Scope: "report", Level: permission.Editor},
			{Scope: "role", Level: permission.Viewer},
			{Scope: "sync", Level: permission.Editor},
			{Scope: "transaction", Level: permission.Editor},
			{Scope: "user", Level: permission.Editor},
		},
	}
	id, err := uc.Create(ctx, request)
	boot.TerminateWhenError(err, request)

	return id
}

func CreateManagerRole(uc role.Usecase) uint {
	ctx := context.TODO()

	role := "manager"

	filter := &request.Filter{Name: &role}
	list, err := uc.List(ctx, filter)
	if errs.Not(err, "not found") {
		boot.TerminateWhenError(err, filter)
	}
	if err == nil && len(list) > 0 {
		return list[0].ID
	}

	log.Info().Msg("create manager role")
	request := &request.Create{
		Name: role,
		Permissions: []request.Permission{
			{Scope: "machine", Level: permission.Owner},
			{Scope: "machine_slot", Level: permission.Owner},
			{Scope: "payment_channel", Level: permission.Owner},
			{Scope: "product", Level: permission.Owner},
			{Scope: "report", Level: permission.Owner},
			{Scope: "role", Level: permission.Viewer},
			{Scope: "sync", Level: permission.Owner},
			{Scope: "transaction", Level: permission.Owner},
			{Scope: "user", Level: permission.Owner},
		},
	}
	id, err := uc.Create(ctx, request)
	boot.TerminateWhenError(err, request)

	return id
}

func CreateStaffRole(uc role.Usecase) uint {
	ctx := context.TODO()

	role := "staff"

	filter := &request.Filter{Name: &role}
	list, err := uc.List(ctx, filter)
	if errs.Not(err, "not found") {
		boot.TerminateWhenError(err, filter)
	}
	if err == nil && len(list) > 0 {
		return list[0].ID
	}

	log.Info().Msg("create staff role")
	request := &request.Create{
		Name: role,
		Permissions: []request.Permission{
			{Scope: "machine", Level: permission.Owner},
			{Scope: "machine_slot", Level: permission.Owner},
			{Scope: "payment_channel", Level: permission.Viewer},
			{Scope: "product", Level: permission.Owner},
			{Scope: "report", Level: permission.Forbidder},
			{Scope: "role", Level: permission.Viewer},
			{Scope: "sync", Level: permission.Owner},
			{Scope: "transaction", Level: permission.Viewer},
			{Scope: "user", Level: permission.Viewer},
		},
	}
	id, err := uc.Create(ctx, request)
	boot.TerminateWhenError(err, request)

	return id
}
