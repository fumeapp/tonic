package tonic

import (
	"github.com/fumeapp/tonic/database"
	"github.com/fumeapp/tonic/render"
	"github.com/fumeapp/tonic/route"
	"github.com/fumeapp/tonic/setting"
	"github.com/gofiber/fiber/v2"
)

func Init(config *fiber.Config) *fiber.App {

	setting.Setup()

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
