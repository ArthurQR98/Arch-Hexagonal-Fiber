package boostrap

import (
	"database/sql"
	"fmt"

	"github.com/ArthurQR98/challenge_fiber/internal/creating"
	"github.com/ArthurQR98/challenge_fiber/internal/platform/bus/inmemory"
	"github.com/ArthurQR98/challenge_fiber/internal/platform/server"
	"github.com/ArthurQR98/challenge_fiber/internal/platform/storage/mysql"
	_ "github.com/go-sql-driver/mysql" //important
)

const (
	host = "localhost"
	port = 3000

	dbUser = "arthur"
	dbPass = "020398"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "test"
)

func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	var (
		commandBus = inmemory.NewCommandBus()
	)

	courseRepository := mysql.NewCourseRepository(db)

	creatingCourseService := creating.NewCourseService(courseRepository)

	createCourseCommandHandler := creating.NewCourseCommandHandler(creatingCourseService)
	commandBus.Register(creating.CourseCommandType, createCourseCommandHandler)

	srv := server.New(host, port, commandBus)
	return srv.Run()
}
