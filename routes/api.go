package routes

import (
	UserController "github.com/fumeapp/skele/controllers"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {

	r := gin.New()

	r.GET("/user", UserController.Index)
	r.GET("/user/:id", UserController.Show)

	return r
}