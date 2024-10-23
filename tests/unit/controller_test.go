package unit

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"

	"github.com/RianAsmara/personal-finance-advisor-api/configuration"
	"github.com/RianAsmara/personal-finance-advisor-api/controller"
	"github.com/RianAsmara/personal-finance-advisor-api/entity"
	"github.com/RianAsmara/personal-finance-advisor-api/exception"
	"github.com/RianAsmara/personal-finance-advisor-api/model"
	repository "github.com/RianAsmara/personal-finance-advisor-api/repository/impl"
	services "github.com/RianAsmara/personal-finance-advisor-api/service/impl"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"golang.org/x/crypto/bcrypt"
)

func createTestApp() *fiber.App {
	// setup fiber
	app := fiber.New(configuration.NewFiberConfiguration())
	app.Use(recover.New())
	app.Use(cors.New())

	// setup router
	userController.Route(app)
	roleController.Route(app)
	authController.Route(app)

	return app
}

// configuration
var config = configuration.New("../../.env.test")
var database = configuration.NewDatabase(config)
var redis = configuration.NewRedis(config)

// repository
var userRepository = repository.NewUserRepositoryImpl(database)
var roleRepository = repository.NewRoleRepositoryImpl(database)
var authRepository = repository.NewAuthRepositoryImpl(database)

// service
var userService = services.NewUserServiceImpl(&userRepository)
var roleService = services.NewRoleServiceImpl(&roleRepository)
var authService = services.NewAuthServiceImpl(&authRepository)

// controller
var userController = controller.NewUserController(&userService, config)
var roleController = controller.NewRoleController(&roleService, config)
var authController = controller.NewAuthController(&authService, config)

var appTest = createTestApp()

func AuthenticationTest() map[string]interface{} {
	// userRepository.DeleteAll()

	password, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	if err != nil {
		exception.PanicLogging(err)
	}

	roles := []entity.Role{
		{Name: "admin"},
		{Name: "user"},
	}

	userRepository.Insert(context.Background(), entity.User{
		Email:    "admin@gmail.com",
		Password: string(password),
		Roles:    roles,
		IsActive: true,
		// FamilyTreeID: uuid.New(),
	})
	if err != nil {
		exception.PanicLogging(err)
	}

	loginRequestModel := model.LoginRequest{
		Email:    "admin@gmail.com",
		Password: "admin",
	}

	loginRequestBody, err := json.Marshal(loginRequestModel)
	if err != nil {
		exception.PanicLogging(err)
	}

	loginRequest := httptest.NewRequest("POST", "/v1/api/auth/login", bytes.NewBuffer(loginRequestBody))
	loginRequest.Header.Set("Content-Type", "application/json")
	loginRequest.Header.Set("Accept", "application/json")

	loginResponse, err := appTest.Test(loginRequest)
	if err != nil {
		exception.PanicLogging(err)
	}
	defer loginResponse.Body.Close()

	loginResponseBody, err := io.ReadAll(loginResponse.Body)
	if err != nil {
		exception.PanicLogging(err)
	}

	var loginWebResponse model.GeneralResponse
	if err := json.Unmarshal(loginResponseBody, &loginWebResponse); err != nil {
		exception.PanicLogging(err)
	}

	loginJsonData, err := json.Marshal(loginWebResponse.Data)
	if err != nil {
		exception.PanicLogging(err)
	}

	var tokenResponse map[string]interface{}
	if err := json.Unmarshal(loginJsonData, &tokenResponse); err != nil {
		exception.PanicLogging(err)
	}

	fmt.Println(tokenResponse)

	return tokenResponse
}
