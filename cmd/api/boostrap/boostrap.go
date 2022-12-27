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
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	// Server configuration
	Host            string        `default:"localhost"`
	Port            uint          `default:"4000"`
	ShutdownTimeout time.Duration `default:"10s"`
	// Database configuration
	DbUser    string        `default:"test"`
	DbPass    string        `default:"test"`
	DbHost    string        `default:"localhost"`
	DbPort    uint          `default:"3306"`
	DbName    string        `default:"test"`
	DbTimeout time.Duration `default:"5s"`
}

func Run() error {
	var cfg config
	err := envconfig.Process("FIBER", &cfg)
	if err != nil {
		return err
	}

	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	var (
		commandBus = inmemory.NewCommandBus()
		eventBus   = inmemory.NewEventBus()
	)

	courseRepository := mysql.NewCourseRepository(db, cfg.DbTimeout)

	creatingCourseService := creating.NewCourseService(courseRepository, eventBus)
	increasingCourseCounterService := increasing.NewCourseCounterService()

	createCourseCommandHandler := creating.NewCourseCommandHandler(creatingCourseService)
	commandBus.Register(creating.CourseCommandType, createCourseCommandHandler)

	eventBus.Subscribe(
		mooc.CourseCreatedEventType,
		creating.NewIncreaseCoursesCounterOnCourseCreated(increasingCourseCounterService),
	)

	srv := server.New(cfg.Host, cfg.Port, cfg.ShutdownTimeout, commandBus)
	return srv.Run()
}
