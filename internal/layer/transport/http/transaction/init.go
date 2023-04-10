package transaction

import (
	"fmt"

	"github.com/aff-vending-machine/vm-backend/internal/layer/usecase"
	"github.com/gofiber/fiber/v2"
)

type httpImpl struct {
	usecase usecase.Transaction
}

func New(uc usecase.Transaction) *httpImpl {
	return &httpImpl{uc}
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
