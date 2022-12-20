package server

import (
	"fmt"

	"github.com/ArthurQR98/challenge_fiber/internal/platform/server/handler/health"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	httpAddr string
	engine   *fiber.App
}

func New(host string, port uint) Server {
	srv := Server{
		engine: fiber.New(fiber.Config{
			CaseSensitive: true,
			StrictRouting: true,
			AppName:       "Challenge Fiber",
		}),
		httpAddr: fmt.Sprintf("%s:%d", host, port),
	}
	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	s.engine.Use(compress.New())
	s.engine.Use(cors.New())
	s.engine.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	return s.engine.Listen(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.Get("/health", health.CheckHandler())
}
