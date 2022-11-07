package templates

var App = `
package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func NewApp() *fiber.App {
	// create app
	app := fiber.New(config.Fiber(false))

	// middleware
	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(logger.New(config.Logger()))

	// Set fiber monitor path
	// https://docs.gofiber.io/api/middleware/monitor
	app.Get("/metrics", monitor.New(config.Monitor()))
}
`
