package render

import (
	"net/http"
	"time"

	"github.com/fumeapp/tonic"
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
	result := gin.H{
		"benchmark": (time.Now().UnixMicro() - tonic.Before),
		"data":      data,
	}
	c.JSON(200, result)
}
