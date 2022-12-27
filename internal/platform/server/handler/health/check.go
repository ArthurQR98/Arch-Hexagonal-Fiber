package health

import (
	"github.com/gofiber/fiber/v2"
)

// CheckHandler godoc
//
//	@Summary		check status of server
//	@Description	check status of server
//	@Tags			Check
//	@Accept			json
//	@Success		200	{object}	string
//	@Produce		json
//	@Router			/health [get]
func CheckHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).SendString("Everything is Ok!")
	}
}
