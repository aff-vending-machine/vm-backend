package payment_transaction

import (
	"fmt"
	"vm-backend/internal/core/domain/payment"

	"github.com/gofiber/fiber/v2"
)

type transportImpl struct {
	usecase payment.TransactionUsecase
}

func NewTransport(uc payment.TransactionUsecase) *transportImpl {
	return &transportImpl{uc}
}

func getUser(c *fiber.Ctx) string {
	if c.Locals("x-access") == nil {
		return "unknown"
	}

	if str, ok := c.Locals("x-access").(string); ok {
		return str
	}

	return fmt.Sprintf("%v", c.Locals("x-access"))
}
