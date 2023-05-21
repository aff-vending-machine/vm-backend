package migration

import (
	"time"
	"vm-backend/internal/core/domain/account"
	"vm-backend/pkg/errs"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type User struct {
	ID        uint       `json:"id" gorm:"primarykey"`
	RoleID    uint       `json:"role_id"`
	Role      Role       `json:"role"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Username  string     `json:"username" gorm:"uniqueIndex"`
	Password  string     `json:"-"`
	CreatedBy string     `json:"created_by"`
	LastLogin *time.Time `json:"last_login"`
}

type Role struct {
	ID          uint         `json:"id" gorm:"primarykey"`
	Permissions []Permission `json:"permissions"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	Name        string       `json:"name" gorm:"uniqueIndex"`
}

type Permission struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	RoleID    uint      `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Scope     string    `json:"scope"`
	Level     int       `json:"level"`
}

func (e User) TableName() string {
	return "users"
}

func (e Role) TableName() string {
	return "roles"
}

func (e Permission) TableName() string {
	return "permissions"
}

func MigrateUser(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Role{}, &Permission{})

	var users []User
	db.Preload("Role").Preload("Role.Permissions").Find(&users)

	for _, user := range users {
		var accountRole account.Role

		// Check if the role already exists in the new table
		if isNotFound(db.Where("name = ?", user.Role.Name), &accountRole) {
			accountRole = account.Role{
				Name:      user.Role.Name,
				CreatedAt: user.Role.CreatedAt,
				UpdatedAt: user.Role.UpdatedAt,
			}
			db.Create(&accountRole)
			log.Info().Str("name", accountRole.Name).Msg("migrated role")

			for _, permission := range user.Role.Permissions {
				var accountPermission account.Permission

				// Check if the address already exists in the new table
				if isNotFound(db.Where("role_id = ? AND scope = ?", permission.RoleID, permission.Scope), &accountPermission) {
					accountPermission = account.Permission{
						RoleID:    accountRole.ID,
						CreatedAt: permission.CreatedAt,
						UpdatedAt: permission.UpdatedAt,
						Scope:     permission.Scope,
						Level:     permission.Level,
					}
					db.Create(&accountPermission)
					log.Info().Uint("role_id", accountPermission.RoleID).Str("scope", accountPermission.Scope).Msg("migrated permission")
				}
			}

		}

		var accountUser account.User

		// Check if the user already exists in the new table
		if isNotFound(db.Where("username = ?", user.Username), &accountUser) {
			accountUser = account.User{
				RoleID:    accountRole.ID,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
				Username:  user.Username,
				Password:  user.Password,
				Contact:   "",
				CreatedBy: user.CreatedBy,
				LastLogin: user.LastLogin,
				LastToken: "",
			}
			db.Create(&accountUser)
			log.Info().Str("username", accountUser.Username).Msg("migrated user")
		}
	}
}

func isNotFound(db *gorm.DB, data interface{}) bool {
	err := db.First(&data).Error
	if errs.IsErr(err, gorm.ErrRecordNotFound) {
		return true
	}
	if err != nil {
		log.Error().Err(err).Interface("data", data).Err(err).Msg("error while checking if the record exists")
		return false
	}

	return false
}
