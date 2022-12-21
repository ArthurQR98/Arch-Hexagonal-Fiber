package health

import (
	"github.com/gofiber/fiber/v2"
)

func CheckHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).SendString("Everything is Ok!")
	}
}
