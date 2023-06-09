package http

import (
	"bytes"
	"fmt"

	"vm-backend/pkg/helpers/errs"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

// 200 - OK
func OK(ctx *fiber.Ctx, data interface{}) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "done",
		"data":   data,
	})
}

// 200 - File
func SendFile(ctx *fiber.Ctx, filename string, buf *bytes.Buffer) error {
	ctx.Set("Content-Type", "text/csv")
	ctx.Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))

	return ctx.Status(fiber.StatusOK).SendStream(buf)
}

// 204 - NoContent
func NoContent(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":   fiber.StatusNoContent,
		"status": "done",
	})
}

// 201 - Created
func Created(ctx *fiber.Ctx, id uint) error {
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":   fiber.StatusCreated,
		"id":     id,
		"status": "done",
	})
}

func UsecaseError(ctx *fiber.Ctx, err error) error {
	code, msg := translateError(err)
	return ctx.Status(code).JSON(fiber.Map{
		"code":    code,
		"status":  "error",
		"message": msg,
	})
}

// 400 - Bad Request
func BadRequest(ctx *fiber.Ctx, cause error) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"code":    fiber.StatusBadRequest,
		"status":  "error",
		"message": cause.Error(),
	})
}

// 401 - Unauthorized
func Unauthorized(ctx *fiber.Ctx, cause error) error {
	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"code":    fiber.StatusUnauthorized,
		"status":  "error",
		"message": cause.Error(),
	})
}

// 403 - Forbidden
func Forbidden(ctx *fiber.Ctx, cause error) error {
	return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
		"code":    fiber.StatusForbidden,
		"status":  "error",
		"message": cause.Error(),
	})
}

func translateError(err error) (int, string) {
	// 400
	if errs.HasMsg(err, "invalid request") {
		return fiber.StatusBadRequest, errors.Cause(err).Error()
	}

	if errs.HasMsg(err, "exist") {
		return fiber.StatusBadRequest, errors.Cause(err).Error()
	}

	if errs.HasMsg(err, "device id") {
		return fiber.StatusBadRequest, "device ID is invalid"
	}

	if errs.HasMsg(err, "decrypt") {
		return fiber.StatusBadRequest, "data is invalid"
	}

	if errs.HasMsg(err, "invalid data") {
		return fiber.StatusBadRequest, "data is invalid"
	}

	// 401
	if errs.HasMsg(err, "password") {
		return fiber.StatusBadRequest, "password is not match"
	}

	// 403
	if errs.HasMsg(err, "signature") {
		return fiber.StatusBadRequest, errors.Cause(err).Error()
	}

	if errors.Is(err, fiber.ErrForbidden) {
		return fiber.StatusForbidden, "no permission"
	}

	if errs.HasMsg(err, "no permission") {
		return fiber.StatusForbidden, err.Error()
	}

	// 404

	// 500
	if errs.HasMsg(err, "rpc error") {
		return fiber.StatusInternalServerError, errors.Cause(err).Error()
	}

	return fiber.StatusBadRequest, fmt.Sprintf("unexpected error: (%s)", err.Error())
}
