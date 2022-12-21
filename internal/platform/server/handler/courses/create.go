package courses

import (
	"errors"

	mooc "github.com/ArthurQR98/challenge_fiber/internal"
	"github.com/ArthurQR98/challenge_fiber/internal/creating"
	"github.com/gofiber/fiber/v2"
)

type createRequest struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

func CreateHandler(creatingCourseService creating.CourseService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		req := new(createRequest)
		if err := ctx.BodyParser(req); err != nil {
			return err
		}

		err := creatingCourseService.CreateCourse(ctx.Context(), req.ID, req.Name, req.Duration)

		if err != nil {
			switch {
			case errors.Is(err, mooc.ErrInvalidCourseID),
				errors.Is(err, mooc.ErrEmptyCourseName),
				errors.Is(err, mooc.ErrEmptyDuration):
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
			default:
				return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
			}
		}

		return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "create successfully",
		})
	}
}
