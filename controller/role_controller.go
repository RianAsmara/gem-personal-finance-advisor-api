package controller

import (
	"github.com/RianAsmara/personal-finance-advisor-api/configuration"
	"github.com/RianAsmara/personal-finance-advisor-api/middleware"
	"github.com/RianAsmara/personal-finance-advisor-api/model"
	services "github.com/RianAsmara/personal-finance-advisor-api/service"
	"github.com/gofiber/fiber/v2"
)

type RoleController struct {
	services.RoleService
	configuration.Config
}

func NewRoleController(roleService *services.RoleService, config configuration.Config) *RoleController {
	return &RoleController{RoleService: *roleService, Config: config}
}

func (controller RoleController) Route(app *fiber.App) {
	app.Get("/v1/api/roles", middleware.AuthenticateJWT("admin", controller.Config), controller.GetRoles)
	app.Get("/v1/api/role/:id", middleware.AuthenticateJWT("admin", controller.Config), controller.GetRole)
	app.Post("/v1/api/role", middleware.AuthenticateJWT("admin", controller.Config), controller.CreateRole)
	app.Put("/v1/api/role/:id", middleware.AuthenticateJWT("admin", controller.Config), controller.UpdateRole)
	app.Delete("/v1/api/role/:id", middleware.AuthenticateJWT("admin", controller.Config), controller.DeleteRole)
}

// GetRoles godoc
// @Summary Get all existing roles
// @Description Get all existing roles
// @Tags Roles
// @Accept json
// @Produce json
// @Success 200 {object} model.GeneralResponse
// @Security ApiKeyAuth
// @Router /v1/api/roles [get]
// @Response 200 {object} model.GeneralResponse{data=[]model.User} "Successful response"
// @Response 401 {object} model.GeneralResponse "Unauthorized"
// @Response 500 {object} model.GeneralResponse "Internal Server Error"
func (controller RoleController) GetRoles(c *fiber.Ctx) error {
	result := controller.RoleService.GetRolesService(c.Context())
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Success: true,
		Message: "Success",
		Data:    result,
	})
}

// GetRoleById godoc
// @Summary Get role by id
// @Description Get role by id
// @Tags Roles
// @Accept json
// @Produce json
// @Param id path string true "Role ID"
// @Success 200 {object} model.GeneralResponse
// @Router /v1/api/role/:id [get]
// @Response 200 {object} model.GeneralResponse{data=model.Role} "Successful response"
// @Response 401 {object} model.GeneralResponse "Unauthorized"
// @Response 500 {object} model.GeneralResponse "Internal Server Error"
func (controller RoleController) GetRole(c *fiber.Ctx) error {
	id := c.Params("id")
	result := controller.RoleService.GetRoleByIdService(c.Context(), id)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Success: true,
		Message: "Success",
		Data:    result,
	})
}

// CreateRole godoc
// @Summary Create a new role
// @Description Create a new role
// @Tags Roles
// @Accept json
// @Produce json
// @Param role body model.RoleRequest true "Role data"
// @Success 200 {object} model.GeneralResponse
// @Router /v1/api/role [post]
// @Response 200 {object} model.GeneralResponse{data=model.RoleRequest} "Successful response"
// @Response 401 {object} model.GeneralResponse "Unauthorized"
// @Response 500 {object} model.GeneralResponse "Internal Server Error"
func (controller RoleController) CreateRole(c *fiber.Ctx) error {
	var request model.RoleRequest
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Success: false,
			Message: "Bad Request",
			Data:    nil,
		})
	}

	result := controller.RoleService.CreateRoleService(c.Context(), request)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Success: true,
		Message: "Success",
		Data:    result,
	})
}

// UpdateRole godoc
// @Summary Update a role
// @Description Update a role
// @Tags Roles
// @Accept json
// @Produce json
// @Param id path string true "Role ID"
// @Param role body model.RoleRequest true "Role data"
// @Success 200 {object} model.GeneralResponse
// @Router /v1/api/role/:id [put]
// @Response 200 {object} model.GeneralResponse{data=model.RoleRequest} "Successful response"
// @Response 401 {object} model.GeneralResponse "Unauthorized"
// @Response 500 {object} model.GeneralResponse "Internal Server Error"
func (controller RoleController) UpdateRole(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Success: false,
			Message: "Bad Request",
			Data:    nil,
		})
	}

	var request model.RoleRequest
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Success: false,
			Message: "Bad Request",
			Data:    nil,
		})
	}

	result := controller.RoleService.UpdateRoleService(c.Context(), request)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Success: true,
		Message: "Success",
		Data:    result,
	})
}

// DeleteRole godoc
// @Summary Delete a role
// @Description Delete a role
// @Tags Roles
// @Accept json
// @Produce json
// @Param id path string true "Role ID"
// @Success 200 {object} model.GeneralResponse
// @Router /v1/api/role/:id [delete]
// @Response 200 {object} model.GeneralResponse "Successful response"
// @Response 401 {object} model.GeneralResponse "Unauthorized"
// @Response 500 {object} model.GeneralResponse "Internal Server Error"
func (controller RoleController) DeleteRole(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Success: false,
			Message: "Bad Request",
			Data:    nil,
		})
	}
	controller.RoleService.DeleteRoleService(c.Context(), id)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Success: true,
		Message: "Success",
		Data:    nil,
	})
}
