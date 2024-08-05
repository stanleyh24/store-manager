package api

import (
	"store-manager/api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewHttp(dbPool *pgxpool.Pool) *fiber.App {
	//userhandler := handlers.NewUserHandler(storage.New(db))

	app := fiber.New()

	routes.InitRoutes(app, dbPool)

	return app
}
