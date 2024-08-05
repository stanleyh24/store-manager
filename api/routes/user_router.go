package routes

import (
	"store-manager/api/handlers"
	"store-manager/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

// NewRouter returns a router to handle model.Invoice requests
func NewUserRouter(app *fiber.App, dbPool *pgxpool.Pool) {

	userHandler := handlers.NewUserHandler(storage.NewUserStorage(dbPool))

	app.Get("/users", userHandler.HandleGetUsers)
	app.Get("/users/roles", userHandler.HandleGetAllRoles)
	app.Get("/users/role/:id", userHandler.HandleGetRole)
	app.Post("/users/role", userHandler.HandleCreateRole)
	app.Put("users/role/:id", userHandler.HandleUpdateRole)
	app.Delete("users/role/:id", userHandler.HandleDeleteRole)

	app.Get("/users/role/operations", userHandler.HandleGetAllOperations)
	app.Post("/users/role/permissions/add", userHandler.HandlerAddRoleOperation)
	app.Post("/users/role/permissions/remove", userHandler.HandlerRemoveRoleOperation)

}
