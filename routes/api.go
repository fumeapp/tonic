package routes

import (
	"github.com/fumeapp/tonic/controllers"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {

	r := gin.New()

	r.GET("/", controllers.BaseIndex)
	r.GET("/user", controllers.UserIndex)
	r.GET("/user/:id", controllers.UserShow)

	return r
}