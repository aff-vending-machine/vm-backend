package report

import "github.com/gofiber/fiber/v2"

type Transport interface {
	Summary(ctx *fiber.Ctx) error      // GET 	{reports/summary}
	Stocks(ctx *fiber.Ctx) error       // GET 	{reports/:id/stocks}
	Transactions(ctx *fiber.Ctx) error // GET 	{reports/:id/transactions}
}
