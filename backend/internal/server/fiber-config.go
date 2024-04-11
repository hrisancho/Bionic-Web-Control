package server

import "github.com/gofiber/fiber/v2"

func fiberErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return ctx.Status(code).JSON(fiber.Map{
		"status": "error",
		"error":  err.Error(),
	})
}

func GetFiberConfig() (config fiber.Config) {
	config.ErrorHandler = fiberErrorHandler
	return
}
