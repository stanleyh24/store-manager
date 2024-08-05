package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitRoutes(app *fiber.App, dbPool *pgxpool.Pool) {
	health(app)

	// A
	// B
	// C

	// I
	//invoice.NewRouter(e, dbPool)

	// L
	//login.NewRouter(e, dbPool)

	// P
	//paypal.NewRouter(e, dbPool)
	//product.NewRouter(e, dbPool)
	//purchaseorder.NewRouter(e, dbPool)

	// U
	NewUserRouter(app, dbPool)
}

func health(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hola")
	})
}
