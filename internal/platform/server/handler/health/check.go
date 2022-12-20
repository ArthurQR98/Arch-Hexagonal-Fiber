package health

import (
	"github.com/gofiber/fiber/v2"
)

func CheckHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ctx.SendString("Everything is Ok!")
		return ctx.SendStatus(fiber.StatusOK)
	}
}
