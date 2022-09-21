package tonic

import (
	"github.com/fumeapp/tonic/database"
	"github.com/fumeapp/tonic/route"
	"github.com/fumeapp/tonic/setting"
	"github.com/gofiber/fiber/v2"
)

func Init(config fiber.Config) *fiber.App {

	setting.Setup()
	database.Setup()

	app := fiber.New(config)
	app.Use(route.Benchmark())
	return app
}
