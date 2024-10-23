package configuration

import (
	"github.com/RianAsmara/personal-finance-advisor-api/exception"
	"github.com/gofiber/fiber/v2"
)

func NewFiberConfiguration() fiber.Config {
	return fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Family Tree API",
		ErrorHandler:  exception.ErrorHandler,
	}
}
