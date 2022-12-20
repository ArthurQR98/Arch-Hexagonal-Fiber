package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

const httpAddr = ":3000"

func main() {
	fmt.Println("Server running on", httpAddr)

	app := fiber.New()
	app.Get("/health", healthHandler)

	app.Listen(httpAddr)
}

func healthHandler(c *fiber.Ctx) error {
	return c.Send([]byte("Everything is ok!"))
}
