package controller

import (
	"github.com/RianAsmara/personal-finance-advisor-api/configuration"
	"github.com/RianAsmara/personal-finance-advisor-api/middleware"
	"github.com/RianAsmara/personal-finance-advisor-api/model"
	services "github.com/RianAsmara/personal-finance-advisor-api/service"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type UserController struct {
	services.UserService
	configuration.Config
	*zap.Logger
}

func NewUserController(userService *services.UserService, config configuration.Config) *UserController {
	return &UserController{UserService: *userService, Config: config}
}

func (controller UserController) Route(app *fiber.App) {
	app.Get("/v1/api/users", middleware.AuthenticateJWT("admin", controller.Config), controller.GetUsers)
	app.Get("/v1/api/user/:id", middleware.AuthenticateJWT("admin", controller.Config), controller.GetUserById)
	app.Post("/v1/api/user", middleware.AuthenticateJWT("admin", controller.Config), controller.CreateUser)
	app.Put("/v1/api/user/:id", middleware.AuthenticateJWT("admin", controller.Config), controller.UpdateUser)
	app.Delete("/v1/api/user/:id", middleware.AuthenticateJWT("admin", controller.Config), controller.DeleteUser)
}

// GetUsers godoc
// @Summary Get all existing users
// @Description Get all existing users
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} model.GeneralResponse
// @Security ApiKeyAuth
// @Router /v1/api/users [get]
// @Response 200 {object} model.GeneralResponse{data=[]model.User} "Successful response"
// @Response 401 {object} model.GeneralResponse "Unauthorized"
// @Response 500 {object} model.GeneralResponse "Internal Server Error"
func (controller UserController) GetUsers(c *fiber.Ctx) error {
	result := controller.UserService.GetUsersService(c.Context())
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Success: true,
		Message: "Success",
		Data:    result,
	})
}

// GetUserById godoc
// @Summary Get user by id
// @Description Get user by id
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} model.GeneralResponse
// @Router /v1/api/user/:id [get]
// @Response 200 {object} model.GeneralResponse{data=model.User} "Successful response"
// @Response 401 {object} model.GeneralResponse "Unauthorized"
// @Response 500 {object} model.GeneralResponse "Internal Server Error"
func (controller UserController) GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	result := controller.UserService.GetUserByIdService(c.Context(), id)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Success: true,
		Message: "Success",
		Data:    result,
	})
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body model.UserRequest true "User data"
// @Success 200 {object} model.GeneralResponse
// @Router /v1/api/user [post]
// @Response 200 {object} model.GeneralResponse{data=model.UserRequest} "Successful response"
// @Response 401 {object} model.GeneralResponse "Unauthorized"
// @Response 500 {object} model.GeneralResponse "Internal Server Error"
func (controller UserController) CreateUser(c *fiber.Ctx) error {
	var request model.UserRequest
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Success: false,
			Message: "Bad Request",
			Data:    nil,
		})
	}

	result := controller.UserService.CreateUserService(c.Context(), request)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Success: true,
		Message: "Success",
		Data:    result,
	})
}

// UpdateUser godoc
// @Summary Update a user
// @Description Update a user
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body model.UserRequest true "User data"
// @Success 200 {object} model.GeneralResponse
// @Router /v1/api/user/:id [put]
// @Response 200 {object} model.GeneralResponse{data=model.UserRequest} "Successful response"
// @Response 401 {object} model.GeneralResponse "Unauthorized"
// @Response 500 {object} model.GeneralResponse "Internal Server Error"
func (controller UserController) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Success: false,
			Message: "Bad Request",
			Data:    nil,
		})
	}

	var request model.UserRequest
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Success: false,
			Message: "Bad Request",
			Data:    nil,
		})
	}

	result := controller.UserService.UpdateUserService(c.Context(), request)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Success: true,
		Message: "Success",
		Data:    result,
	})
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} model.GeneralResponse
// @Router /v1/api/user/:id [delete]
// @Response 200 {object} model.GeneralResponse "Successful response"
// @Response 401 {object} model.GeneralResponse "Unauthorized"
// @Response 500 {object} model.GeneralResponse "Internal Server Error"
func (controller UserController) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Success: false,
			Message: "Bad Request",
			Data:    nil,
		})
	}
	controller.UserService.DeleteUserService(c.Context(), id)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Success: true,
		Message: "Success",
		Data:    nil,
	})
}
