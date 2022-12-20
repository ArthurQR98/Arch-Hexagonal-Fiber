package courses

import (
	mooc "github.com/ArthurQR98/challenge_fiber/internal"
	_ "github.com/go-sql-driver/mysql" //important
	"github.com/gofiber/fiber/v2"
)

type createRequest struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

func CreateHandler(courseRepository mooc.CourseRepository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		req := new(createRequest)
		if err := ctx.BodyParser(req); err != nil {
			return err
		}
		course, err := mooc.NewCourse(req.ID, req.Name, req.Duration)
		if err != nil {
			ctx.SendStatus(400)
			return ctx.JSON(fiber.Map{"error": err.Error()})
		}
		if err := courseRepository.Save(ctx.Context(), course); err != nil {
			ctx.SendStatus(500)
			return ctx.JSON(fiber.Map{"error": err.Error()})
		}
		ctx.JSON(fiber.Map{
			"message": "create successfully",
		})
		return ctx.SendStatus(fiber.StatusCreated)
	}
}
