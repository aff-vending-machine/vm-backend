package account

import (
	"context"
	"fmt"
	"strconv"
	"vm-backend/internal/layer/usecase/account/request"
	"vm-backend/internal/layer/usecase/account/response"

	"github.com/gofiber/fiber/v2"
)

const (
	PermissionLevelKey = "x-permission-level"
	AccessScopeKey     = "x-access-scope"
	AccessConditionKey = "x-access-condition"
	AccessIDKey        = "x-access-id"
	AccessKey          = "x-access"
	RoleIDKey          = "x-role-id"
	RoleKey            = "x-role"
	BranchIDKey        = "x-branch-id"
	BranchKey          = "x-branch"
)

type Usecase interface {
	Login(ctx context.Context, req *request.Login) (*response.AuthResult, error)
	ValidateAccessToken(ctx context.Context, req *request.ValidateToken) (*response.Claims, error)
	ValidateRefreshToken(ctx context.Context, req *request.ValidateToken) (*response.AuthResult, error)
	GetAccountPermission(ctx context.Context, req *request.GetAccountPermission) (*response.AccountPermission, error)
}

type PasswordManagement interface {
	HashPassword(ctx context.Context, password string) (string, error)
	Compare(ctx context.Context, hashed string, password string) (bool, error)
}

type TokenManagement interface {
	CreateAccessToken(context.Context, Token) (string, error)
	CreateRefreshToken(context.Context, Token) (string, error)
	ValidateAccessToken(context.Context, string) (*Token, error)
	ValidateRefreshToken(context.Context, string) (*Token, error)
}

type Transport interface {
	Login(ctx *fiber.Ctx) error                 // POST {auth/login}
	RefreshToken(ctx *fiber.Ctx) error          // POST {auth/refresh-token}
	AuthorizationRequired(ctx *fiber.Ctx) error // middleware
	PermissionRequired(ctx *fiber.Ctx) error    // middleware
}

func GetAccessID(ctx *fiber.Ctx) (uint, error) {
	if ctx.Locals(AccessIDKey) == nil {
		return 0, fmt.Errorf("user id is not exist")
	}

	str := fmt.Sprintf("%v", ctx.Locals(AccessIDKey))
	id, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}

	return uint(id), nil
}

func GetAccess(c *fiber.Ctx) string {
	if c.Locals(AccessKey) == nil {
		return "unknown"
	}

	if str, ok := c.Locals(AccessKey).(string); ok {
		return str
	}

	return fmt.Sprintf("%v", c.Locals(AccessKey))
}

func GetBranchID(c *fiber.Ctx, id *uint) *uint {
	rejected := uint(0)

	if c.Locals(PermissionLevelKey) == nil {
		return &rejected
	}

	if permissionLevel, ok := c.Locals(PermissionLevelKey).(int); ok {
		if permissionLevel >= Editor {
			// use from filter
			return id
		}
	}

	if c.Locals(BranchIDKey) == nil {
		return &rejected
	}

	if branchID, ok := c.Locals(BranchIDKey).(uint); ok {
		if branchID == 0 {
			// use from filter
			return id
		}

		// use from account
		return &branchID
	}

	return &rejected
}
