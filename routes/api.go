package routes

import (
	usercontroller "github.com/fumeapp/tonic/controllers"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {

	r := gin.New()

	r.GET("/user", usercontroller.Index)
	r.GET("/user/:id", usercontroller.Show)

	return r
}