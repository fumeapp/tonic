package render

import "github.com/gin-gonic/gin"

func Error(c *gin.Context, message string) {
	c.JSON(500, gin.H{"error": message})
	return
}

func Render(c *gin.Context, data gin.H) {
	c.JSON(200, data)
	return
}
