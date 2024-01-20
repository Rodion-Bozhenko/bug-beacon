// Package main is an entrypoint to bug-beacon server
package main

import (
	"bug-beacon/components"
	"log/slog"
	"os"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func main() {
	app := fiber.New()

	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{AddSource: true}))

	app.Static("/assets", "./assets")
	app.Get("/", func(c *fiber.Ctx) error {
		return Render(c, components.Home())
	})

	if err := app.Listen("127.0.0.1:8080"); err != nil {
		logger.Error("Failed to start server", "err", err)
	}
}

// Render renders passed templ component
func Render(c *fiber.Ctx, component templ.Component, options ...func(*templ.ComponentHandler)) error {
	componentHandler := templ.Handler(component)
	for _, o := range options {
		o(componentHandler)
	}
	return adaptor.HTTPHandler(componentHandler)(c)
}
