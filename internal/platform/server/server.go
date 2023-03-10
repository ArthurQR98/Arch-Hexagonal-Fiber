package server

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ArthurQR98/challenge_fiber/internal/platform/server/handler/courses"
	"github.com/ArthurQR98/challenge_fiber/internal/platform/server/handler/health"
	"github.com/ArthurQR98/challenge_fiber/internal/platform/server/middlewares/notfound"
	"github.com/ArthurQR98/challenge_fiber/kit/command"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

type Server struct {
	httpAddr string
	engine   *fiber.App

	// dependencies
	commandBus command.Bus
}

func New(host string, port uint, shutdownTimeout time.Duration, commandBus command.Bus) Server {
	srv := Server{
		engine: fiber.New(fiber.Config{
			CaseSensitive: true,
			StrictRouting: true,
			AppName:       "Challenge Fiber",
			IdleTimeout:   shutdownTimeout,
		}),
		httpAddr:   fmt.Sprintf("%s:%d", host, port),
		commandBus: commandBus,
	}
	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	go func() {
		if err := s.engine.Listen(s.httpAddr); err != nil {
			log.Panic(err)
		}
	}()
	// Create channel to signify a signal being sent
	c := make(chan os.Signal, 1)
	// When an interrupt or termination signal is sent, notify the channel
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	// This blocks the main thread until an interrupt is received
	<-c
	return s.engine.Shutdown()
}

func (s *Server) registerRoutes() {
	// Middlewares
	s.engine.Use(logger.New(logger.Config{
		Format:     "${cyan}[${time}] ${white}${pid} ${red}${status} ${blue}[${method}] ${white}${path}\n",
		TimeFormat: "02-Jan-2006",
	}))
	s.engine.Use(compress.New())
	s.engine.Use(cors.New())
	s.engine.Use(recover.New())

	s.engine.Get("/api-docs/*", swagger.HandlerDefault)
	s.engine.Get("/health", health.CheckHandler())
	s.engine.Post("/courses", courses.CreateHandler(s.commandBus))

	s.engine.Use(notfound.New())
}
