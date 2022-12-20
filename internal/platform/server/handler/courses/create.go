package courses

import (
	"database/sql"
	"fmt"

	mooc "github.com/ArthurQR98/challenge_fiber/internal"
	"github.com/ArthurQR98/challenge_fiber/internal/platform/storage/mysql"
	_ "github.com/go-sql-driver/mysql" //important
	"github.com/gofiber/fiber/v2"
)

const (
	dbUser = "arthur"
	dbPass = "020398"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "test"
)

type createRequest struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

func CreateHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		req := new(createRequest)
		if err := ctx.BodyParser(req); err != nil {
			return err
		}
		course := mooc.NewCourse(req.ID, req.Name, req.Duration)
		mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
		db, err := sql.Open("mysql", mysqlURI)
		if err != nil {
			ctx.JSON(err.Error())
			return err
		}
		courseRepository := mysql.NewCourseRepository(db)
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
