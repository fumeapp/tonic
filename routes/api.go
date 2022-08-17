package routes

import (
	"fmt"

	"github.com/fumeapp/tonic/controllers"
	"github.com/fumeapp/tonic/pkg/route"
	"github.com/gin-gonic/gin"
)

func Init(IsDev bool) *gin.Engine {
	r := gin.New()
	if (IsDev) {
		route.Base(r)
	}
	route.ApiResource(r, "user", controllers.UserResource())
	r.NoRoute(func(c *gin.Context) {
		fmt.Println("Method: " + c.Request.Method)
		fmt.Println("URL: " + c.Request.URL.String())
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	return r
}