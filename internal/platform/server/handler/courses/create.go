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
		course := mooc.NewCourse(req.ID, req.Name, req.Duration)
		if err := courseRepository.Save(ctx.Context(), course); err != nil {
			ctx.JSON(err.Error())
			return err
		}
		ctx.JSON(fiber.Map{
			"message": "create successfully",
		})
		return ctx.SendStatus(fiber.StatusCreated)
	}
}
