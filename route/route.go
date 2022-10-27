package route

import (
	"errors"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"time"

	"github.com/fumeapp/tonic/database"
	"github.com/gofiber/fiber/v2"
)

type ApiResourceStruct struct {
	Index  func(c *fiber.Ctx) error
	Create func(c *fiber.Ctx) error
	Show   func(c *fiber.Ctx, value any) error
	Update func(c *fiber.Ctx, value any) error
	Delete func(c *fiber.Ctx, value any) error
}

var (
	model     any
	resources ApiResourceStruct
	UUID      uuid.UUID
)

type binder func(c *fiber.Ctx, value any) error

func Benchmark(c *fiber.Ctx) error {
	c.Locals("tonicBenchmark", time.Now().UnixMicro())
	return c.Next()
}

func GenerateUUID() {
	UUID = uuid.New()
}

func UUIDMiddleware(c *fiber.Ctx) error {
	c.Locals("tonicUUID", UUID.String())
	return c.Next()
}

func bind(c *fiber.Ctx, callback binder) error {
	if isNumeric(c) {
		value, err := retrieve(c)
		if err != nil {
			return invalid(c)
		} else {
			return callback(c, value)
		}
	} else {
		return invalid(c)
	}
}

func bindShow(c *fiber.Ctx) error {
	return bind(c, resources.Show)
}

func bindUpdate(c *fiber.Ctx) error {
	return bind(c, resources.Update)
}

func bindDelete(c *fiber.Ctx) error {
	return bind(c, resources.Delete)
}

func ApiResource(app *fiber.App, n string, _model any, _resources ApiResourceStruct, middleware any) {
	resources = _resources
	model = _model

	if middleware != nil {
		mid := middleware.(fiber.Handler)
		app.Get("/"+n, mid, resources.Index).Name(n + " Index")
		app.Post("/"+n, mid, resources.Create).Name(n + " Create")
		app.Get("/"+n+"/:id", mid, bindShow).Name(n + " Show")
		app.Put("/"+n+"/:id", mid, bindUpdate).Name(n + " Update")
		app.Delete("/"+n+"/:id", mid, bindDelete).Name(n + " Delete")
	} else {
		app.Get("/"+n, resources.Index).Name(n + " Index")
		app.Post("/"+n, resources.Create).Name(n + " Create")
		app.Get("/"+n+"/:id", bindShow).Name(n + " Show")
		app.Put("/"+n+"/:id", bindUpdate).Name(n + " Update")
		app.Delete("/"+n+"/:id", bindDelete).Name(n + " Delete")
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
