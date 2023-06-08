package tonic

import (
	"github.com/fumeapp/tonic/database"
	"github.com/fumeapp/tonic/render"
	"github.com/fumeapp/tonic/route"
	"github.com/fumeapp/tonic/setting"
	"github.com/gofiber/fiber/v2"
)

func Init(config *fiber.Config, args ...string) *fiber.App {

	if len(args) > 0 {
		setting.Setup(args[0])
	} else {
		setting.Setup(".env")
	}

	config.EnablePrintRoutes = setting.IsDev()

	app := fiber.New(*config)
	app.Use(route.Benchmark)
	if setting.IsDev() {
		app.Get("/", route.List).Name("Route List")
	}

	database.Setup()

	return app
}

func ShowUUID(app *fiber.App) {
	render.ShowUUID = true
	route.GenerateUUID()
	app.Use(route.UUIDMiddleware)
}
