package main

import (
	"github.com/fumeapp/skele/models"
	"github.com/fumeapp/skele/pkg/setting"
	"github.com/gin-gonic/gin"
)


func init() {
	setting.Setup()
	models.Setup()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}