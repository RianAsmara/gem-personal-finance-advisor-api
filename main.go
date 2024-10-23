package main

import (
	"github.com/RianAsmara/personal-finance-advisor-api/client/restclient"
	"github.com/RianAsmara/personal-finance-advisor-api/configuration"
	"github.com/RianAsmara/personal-finance-advisor-api/controller"
	"github.com/RianAsmara/personal-finance-advisor-api/exception"
	"github.com/RianAsmara/personal-finance-advisor-api/middleware"
	repository "github.com/RianAsmara/personal-finance-advisor-api/repository/impl"
	service "github.com/RianAsmara/personal-finance-advisor-api/service/impl"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"github.com/gofiber/swagger"

	_ "github.com/RianAsmara/personal-finance-advisor-api/docs"
)

// @title Genealogy API
// @version 1.0
// @description [ Base URL : http://localhost:9999/v1/api ]
// @description This API using combination JWT & HMAC Authentication
// @description Generate Here https://www.devglan.com/online-tools/hmac-sha256-online
// @description Convert UNIXMilis Here https://www.unixtimestamp.com/
// @schemes http https
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	config := configuration.New()
	database := configuration.NewDatabase(config)
	// redis := configuration.NewRedis(config)
	// newRelicConfig := configuration.NewRelicConfig(config)
	// zapLogger := configuration.ZapConfig(newRelicConfig)

	httpBinRestClient := restclient.NewHttpBinRestClient()

	userRepository := repository.NewUserRepositoryImpl(database)
	authRepository := repository.NewAuthRepositoryImpl(database)

	httpBinService := service.NewHttpBinServiceImpl(&httpBinRestClient)
	userService := service.NewUserServiceImpl(&userRepository)
	authService := service.NewAuthServiceImpl(&authRepository)

	httpBinController := controller.NewHttpBinController(&httpBinService)
	userController := controller.NewUserController(&userService, config)
	authController := controller.NewAuthController(&authService, config)

	app := fiber.New(configuration.NewFiberConfiguration())
	app.Use(healthcheck.New())
	app.Use(recover.New())
	app.Use(cors.New())

	// app.Use(idempotency.New(idempotency.Config{
	// 	Lifetime: 42 * time.Minute,
	// }))
	// Limiter
	// app.Use(limiter.New(limiter.Config{
	// 	Max:               20,
	// 	Expiration:        30 * time.Second,
	// 	LimiterMiddleware: limiter.SlidingWindow{},
	// }))

	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
	}))

	httpBinController.Route(app)
	userController.Route(app)
	authController.Route(app)

	app.Get("/docs/*", middleware.BasicAuthMiddleware("admin", "admin"), swagger.HandlerDefault)

	app.Get("/metrics", monitor.New(monitor.Config{Title: "Family Tree API Metrics Page"}))

	err := app.Listen(config.Get("PORT"))
	exception.PanicLogging(err)
}
