package notfound

import "github.com/gofiber/fiber/v2"

func New() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Resource not found!"})
	}
}
