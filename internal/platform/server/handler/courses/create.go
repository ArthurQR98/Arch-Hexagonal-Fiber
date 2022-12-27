package courses

import (
	"errors"

	mooc "github.com/ArthurQR98/challenge_fiber/internal"
	"github.com/ArthurQR98/challenge_fiber/internal/creating"
	"github.com/ArthurQR98/challenge_fiber/kit/command"
	"github.com/gofiber/fiber/v2"
)

// Course model info
//	@Description	Course model information
//	@Description	with id, name and duration
type createRequest struct {
	ID       string `json:"id" binding:"required"`       // ID this is id uuid
	Name     string `json:"name" binding:"required"`     // This is name
	Duration string `json:"duration" binding:"required"` // this is duration
}

// CreateHandler godoc
//
//	@Summary		create course
//	@Description	create course
//	@Tags			Courses
//	@Accept			json
//	@Produce		json
//	@Param			course	body		createRequest	true	"Add course"
//	@Success		200		{string}	string
//	@Failure		400		{string}	string	"error"
//	@Failure		404		{object}	string
//	@Failure		500		{string}	string	"error"
//	@Router			/courses [post]
func CreateHandler(commandBus command.Bus) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var req createRequest
		if err := ctx.BodyParser(&req); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		err := commandBus.Dispatch(ctx.Context(), creating.NewCourseCommand(req.ID, req.Name, req.Duration))
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
