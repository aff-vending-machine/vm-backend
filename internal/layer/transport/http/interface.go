package http

import "github.com/gofiber/fiber/v2"

type Auth interface {
	Login(ctx *fiber.Ctx) error
	RefreshToken(ctx *fiber.Ctx) error
	AuthorizationRequired(ctx *fiber.Ctx) error
	PermissionRequired(ctx *fiber.Ctx) error
}

type Machine interface {
	Read(ctx *fiber.Ctx) error    // GET
	Count(ctx *fiber.Ctx) error   // GET {count}
	ReadOne(ctx *fiber.Ctx) error // GET {:id}
	Create(ctx *fiber.Ctx) error  // POST
	Update(ctx *fiber.Ctx) error  // PUT {:id}
	Delete(ctx *fiber.Ctx) error  // DELETE {:id}
}

type MachineSlot interface {
	Read(ctx *fiber.Ctx) error    // GET 		{machines/:machine_id/slots}
	Count(ctx *fiber.Ctx) error   // GET 		{machines/:machine_id/slots/count}
	ReadOne(ctx *fiber.Ctx) error // GET 		{machines/:machine_id/slots/:id}
	Create(ctx *fiber.Ctx) error  // POST 	{machines/:machine_id/slots}
	Update(ctx *fiber.Ctx) error  // PUT 		{machines/:machine_id/slots/:id}
	Delete(ctx *fiber.Ctx) error  // DELETE {machines/:machine_id/slots/:id}
	SyncGet(ctx *fiber.Ctx) error // POST 	{machines/:machine_id/slots/sync-get}
	SyncSet(ctx *fiber.Ctx) error // POST 	{machines/:machine_id/slots/sync-set}
}

type PaymentChannel interface {
	Read(ctx *fiber.Ctx) error    // GET
	Count(ctx *fiber.Ctx) error   // GET {count}
	ReadOne(ctx *fiber.Ctx) error // GET {:id}
	Create(ctx *fiber.Ctx) error  // POST
	Active(ctx *fiber.Ctx) error  // POST {:id/active}
	Update(ctx *fiber.Ctx) error  // PUT {:id}
	Delete(ctx *fiber.Ctx) error  // DELETE {:id}
}

type Product interface {
	Read(ctx *fiber.Ctx) error    // GET 		{products}
	Count(ctx *fiber.Ctx) error   // GET 		{products/count}
	ReadOne(ctx *fiber.Ctx) error // GET 		{products/:id}
	Create(ctx *fiber.Ctx) error  // POST 	{products}
	Update(ctx *fiber.Ctx) error  // PUT 		{products/:id}
	Delete(ctx *fiber.Ctx) error  // DELETE {products/:id}
}

type Role interface {
	Read(ctx *fiber.Ctx) error    // GET
	Count(ctx *fiber.Ctx) error   // GET {count}
	ReadOne(ctx *fiber.Ctx) error // GET {:id}
	Create(ctx *fiber.Ctx) error  // POST
	Update(ctx *fiber.Ctx) error  // PUT {:id}
	Delete(ctx *fiber.Ctx) error  // DELETE {:id}
}

type Transaction interface {
	Read(ctx *fiber.Ctx) error    // GET		{transactions}
	Count(ctx *fiber.Ctx) error   // GET 		{transactions/count}
	ReadOne(ctx *fiber.Ctx) error // GET 		{transactions/:id}
	Create(ctx *fiber.Ctx) error  // POST 	{transactions}
	Update(ctx *fiber.Ctx) error  // PUT 		{transactions/:id}
	Delete(ctx *fiber.Ctx) error  // DELETE {transactions/:id}
}

type User interface {
	Read(ctx *fiber.Ctx) error           // GET
	Count(ctx *fiber.Ctx) error          // GET {count}
	ReadOne(ctx *fiber.Ctx) error        // GET {:id}
	Create(ctx *fiber.Ctx) error         // POST
	ChangeRole(ctx *fiber.Ctx) error     // POST {:id/change-role}
	ChangePassword(ctx *fiber.Ctx) error // POST {:id/change-password}
	ResetPassword(ctx *fiber.Ctx) error  // POST {:id/reset-password}
	// Update(ctx *fiber.Ctx) error         // PUT {:id}
	Delete(ctx *fiber.Ctx) error // DELETE {:id}
}
