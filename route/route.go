package route

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/fumeapp/tonic/database"
	"github.com/gofiber/fiber/v2"
)

type ApiResourceStruct struct {
	Index  func(c *fiber.Ctx) error
	Show   func(c *fiber.Ctx, value any) error
	Update func(c *fiber.Ctx, value any) error
	Delete func(c *fiber.Ctx, value any) error
}

var (
	model     any
	resources ApiResourceStruct
)

func Benchmark(c *fiber.Ctx) error {
	c.Locals("tonicBenchmark", time.Now().UnixMicro())
	return c.Next()
}

func bind(c *fiber.Ctx) error {
	if isNumeric(c) {
		value, error := retrieve(c)
		if error != nil {
			return invalid(c)
		} else {
			return resources.Show(c, value)
		}
	} else {
		return invalid(c)
	}
}

func ApiResource(app *fiber.App, n string, _model any, _resources ApiResourceStruct, middleware any) {
	resources = _resources
	model = _model

	if middleware != nil {
		mid := middleware.(fiber.Handler)
		app.Get("/"+n, mid, resources.Index).Name(n + " Index")
		app.Get("/"+n+"/:id", mid, bind).Name(n + " Show")
		app.Put("/"+n+"/:id", mid, bind).Name(n + " Update")
		app.Delete("/"+n+"/:id", mid, bind).Name(n + " Delete")
	} else {
		app.Get("/"+n, resources.Index).Name(n + " Index")
		app.Get("/"+n+"/:id", bind).Name(n + " Show")
		app.Put("/"+n+"/:id", bind).Name(n + " Update")
		app.Delete("/"+n+"/:id", bind).Name(n + " Delete")
	}
}

func isNumeric(c *fiber.Ctx) bool {
	if _, err := strconv.Atoi(c.Params("id")); err != nil {
		return false
	}
	return true
}

func retrieve(c *fiber.Ctx) (any, error) {
	result := database.Db.First(&model, c.Params("id"))
	if result.Error != nil {
		return -1, errors.New("Record not found : " + c.Params("id"))
	}
	return model, nil
}

func invalid(c *fiber.Ctx) error {
	return c.Status(http.StatusNotFound).JSON(&fiber.Map{"message": "Resource not found"})
}
