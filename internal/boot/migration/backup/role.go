package migration

import (
	"vm-backend/internal/core/domain/account"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func MigrateRolePermission(db *gorm.DB) {
	var permissions []account.Permission
	db.Find(&permissions)

	for _, permission := range permissions {
		needSave := false
		if permission.Scope == "machine_slot" {
			needSave = true
			permission.Scope = "slot"
		}

		if permission.Scope == "payment_channel" {
			needSave = true
			permission.Scope = "channel"
		}

		if needSave {
			db.Save(&permission)
			log.Info().Uint("role_id", permission.RoleID).Str("scope", permission.Scope).Msg("migrated permission")
		}
	}

	var roles []account.Role
	db.Preload("Permissions").Find(&roles)

	level := account.Forbidder

	for _, role := range roles {
		needSave := false
		switch role.Name {
		case "super-admin":
			level = account.Admin
		case "admin":
			level = account.Editor
		case "manager":
			level = account.Owner
		case "staff":
			level = account.Viewer
		}

		if role.HasPermission("group") == 0 {
			needSave = true
			role.Permissions = append(role.Permissions, account.Permission{Scope: "group", Level: level})
		}

		if role.HasPermission("branch") == 0 {
			needSave = true
			role.Permissions = append(role.Permissions, account.Permission{Scope: "branch", Level: level})
		}

		if needSave {
			db.Save(&role)
			log.Info().Str("role", role.Name).Msg("migrated role")
		}

	}
}
