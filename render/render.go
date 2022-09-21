package render

import (
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

type H map[string]any

func Success(c *fiber.Ctx, message string, data any) error {
	return c.Status(http.StatusAccepted).JSON(H{
		"benmchmark": bench(c),
		"data": gin.H{
			"type":    "success",
			"success": true,
			"message": message,
			"data":    data,
		},
	})
}

func Error(c *fiber.Ctx, errors any) error {
	if reflect.TypeOf(errors).Kind() == reflect.String {
		errors = [1]string{fmt.Sprintf("%v", errors)}
	}
	return c.Status(http.StatusBadRequest).JSON(H{"error": true, "errors": errors})
}

func Render(c *fiber.Ctx, data any) error {
	return c.Status(http.StatusAccepted).JSON(H{
		"benchmark": bench(c),
		"data":      data,
	})
}

func bench(c *fiber.Ctx) float64 {
	benchmark := c.Locals("tonicBenchmark")
	diff := (float64(time.Now().UnixMicro() - benchmark.(int64))) / 1000000
	return diff
}
