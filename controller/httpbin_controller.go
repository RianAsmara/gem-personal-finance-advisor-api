package controller

import (
	"github.com/RianAsmara/personal-finance-advisor-api/model"
	services "github.com/RianAsmara/personal-finance-advisor-api/service"
	"github.com/gofiber/fiber/v2"
)

type HttpBinController struct {
	services.HttpBinService
}

func NewHttpBinController(httpBinService *services.HttpBinService) *HttpBinController {
	return &HttpBinController{HttpBinService: *httpBinService}
}

func (controller HttpBinController) Route(app *fiber.App) {
	app.Get("/v1/api/httpbin", controller.PostHttpBin)
}

func (controller HttpBinController) PostHttpBin(c *fiber.Ctx) error {

	controller.HttpBinService.PostMethod(c.Context())
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Success: true,
		Message: "Success",
		Data:    nil,
	})
}
