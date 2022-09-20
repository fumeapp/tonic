package render

import (
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
)

type H map[string]any

func Success(c *gin.Context, message string, data any) {
	c.JSON(http.StatusAccepted, H{
		"benmchmark": bench(c),
		"data": gin.H{
			"type":    "success",
			"success": true,
			"message": message,
			"data":    data,
		},
	})
}

func Error(c *gin.Context, errors any) {
	if reflect.TypeOf(errors).Kind() == reflect.String {
		errors = [1]string{fmt.Sprintf("%v", errors)}
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, H{"error": true, "errors": errors})
}

func Render(c *gin.Context, data any) {
	c.JSON(200, H{
		"benchmark": bench(c),
		"data":      data,
	})
}

func bench(c *gin.Context) float64 {
	benchmark, _ := c.Get("tonicBenchmark")
	diff := (float64(time.Now().UnixMicro() - benchmark.(int64))) / 1000000
	return diff
}
