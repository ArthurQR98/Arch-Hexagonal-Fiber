package boostrap

import (
	"database/sql"
	"fmt"
	"time"

	mooc "github.com/ArthurQR98/challenge_fiber/internal"
	"github.com/ArthurQR98/challenge_fiber/internal/creating"
	"github.com/ArthurQR98/challenge_fiber/internal/increasing"
	"github.com/ArthurQR98/challenge_fiber/internal/platform/bus/inmemory"
	"github.com/ArthurQR98/challenge_fiber/internal/platform/server"
	"github.com/ArthurQR98/challenge_fiber/internal/platform/storage/mysql"
	_ "github.com/go-sql-driver/mysql" //important
)

const (
	host            = "localhost"
	port            = 3000
	shutdownTimeout = 10 * time.Second

	dbUser    = "arthur"
	dbPass    = "020398"
	dbHost    = "localhost"
	dbPort    = "3306"
	dbName    = "test"
	dbTimeout = 5 * time.Second
)

func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	var (
		commandBus = inmemory.NewCommandBus()
		eventBus   = inmemory.NewEventBus()
	)

	courseRepository := mysql.NewCourseRepository(db, dbTimeout)

	creatingCourseService := creating.NewCourseService(courseRepository, eventBus)
	increasingCourseCounterService := increasing.NewCourseCounterService()

	createCourseCommandHandler := creating.NewCourseCommandHandler(creatingCourseService)
	commandBus.Register(creating.CourseCommandType, createCourseCommandHandler)

	eventBus.Subscribe(
		mooc.CourseCreatedEventType,
		creating.NewIncreaseCoursesCounterOnCourseCreated(increasingCourseCounterService),
	)

	srv := server.New(host, port, shutdownTimeout, commandBus)
	return srv.Run()
}
