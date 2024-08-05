package handlers

import (
	"store-manager/storage"
	"store-manager/types"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userStore storage.UserStore
}

func NewUserHandler(userStore storage.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users, _ := h.userStore.GetAllUsers(c.Context())
	return c.JSON(users)
}

func (h *UserHandler) HandleGetAllRoles(c *fiber.Ctx) error {
	roles, err := h.userStore.GetAllRoles(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(roles)
}

func (h *UserHandler) HandleGetRole(c *fiber.Ctx) error {
	param := c.Params("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	role, err := h.userStore.GetRole(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(role)
}

func (h *UserHandler) HandleCreateRole(c *fiber.Ctx) error {
	var params types.RoleCreateParams
	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	role, err := h.userStore.CreateRole(c.Context(), params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(role)
}

func (h *UserHandler) HandleUpdateRole(c *fiber.Ctx) error {
	var role types.Role
	param := c.Params("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := c.BodyParser(&role); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if role.ID != id {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "id"})
	}

	roleUpdated, err := h.userStore.UpdateRole(c.Context(), role)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(roleUpdated)
}

func (h *UserHandler) HandleDeleteRole(c *fiber.Ctx) error {
	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err = h.userStore.DeleteRole(c.Context(), id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusOK)

}

func (h *UserHandler) HandleGetAllOperations(c *fiber.Ctx) error {
	operations, err := h.userStore.GetAllOperations(c.Context())

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(operations)
}

func (h *UserHandler) HandlerAddRoleOperation(c *fiber.Ctx) error {
	var operations types.PermisioParams
	if err := c.BodyParser(&operations); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.userStore.AddRoleOperation(c.Context(), operations.IDRole, operations.Permissions)

	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *UserHandler) HandlerRemoveRoleOperation(c *fiber.Ctx) error {
	var operations types.PermisioParams
	if err := c.BodyParser(&operations); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.userStore.DeleteRoleOperation(c.Context(), operations.IDRole, operations.Permissions)

	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusOK)
}
