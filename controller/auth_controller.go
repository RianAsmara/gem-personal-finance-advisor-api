package controller

import (
	"github.com/RianAsmara/personal-finance-advisor-api/common"
	"github.com/RianAsmara/personal-finance-advisor-api/configuration"
	"github.com/RianAsmara/personal-finance-advisor-api/exception"
	"github.com/RianAsmara/personal-finance-advisor-api/model"
	services "github.com/RianAsmara/personal-finance-advisor-api/service"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	services.AuthService
	configuration.Config
}

func NewAuthController(authService *services.AuthService, config configuration.Config) *AuthController {
	return &AuthController{AuthService: *authService, Config: config}
}

func (controller AuthController) Route(app *fiber.App) {
	app.Post("/v1/api/auth/login", controller.Login)
}

// Authentication godoc
// @Summary Authenticate user
// @Description Authenticate user and return user information with JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body model.LoginRequest true "User credentials"
// @Success 200 {object} model.GeneralResponse
// @Failure 401 {object} model.GeneralResponse "Unauthorized"
// @Failure 500 {object} model.GeneralResponse "Internal Server Error"
// @Router /v1/api/auth/login [post]
// @Example
//
//	{
//	  "email": "user@example.com",
//	  "password": "password123"
//	}
func (controller AuthController) Login(c *fiber.Ctx) error {
	var request model.LoginRequest

	err := c.BodyParser(&request)

	exception.PanicLogging(err)

	result := controller.AuthService.AuthenticationService(c.Context(), request)
	var roles []map[string]interface{}

	for _, role := range result.Roles {
		roles = append(roles, map[string]interface{}{
			"id":   role.ID,
			"name": role.Name,
		})
	}

	tokenJwt, err := common.GenerateJWT(result.Email, roles, controller.Config)

	exception.PanicLogging(err)

	resultWithToken := map[string]interface{}{
		"email": result.Email,
		"roles": roles,
		"token": tokenJwt,
	}

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Success: true,
		Message: "Success",
		Data:    resultWithToken,
	})
}
