package render

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, message string, data any) {
	c.JSON(http.StatusAccepted, gin.H{
		"success": true,
		"message": message,
		"data":    data,
	})
}

func Error(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{"error": message})
	c.Abort()
}

func Render(c *gin.Context, data any) {
	benchmark, _ := c.Get("tonicBenchmark")
	diff := (float64(time.Now().UnixMicro() - benchmark.(int64))) / 1000000
	result := gin.H{
		"benchmark": diff,
		"data":      data,
	}
	c.JSON(200, result)
}
