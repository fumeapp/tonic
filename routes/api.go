package routes

import (
	"fmt"

	"github.com/fumeapp/tonic/controllers"
	"github.com/fumeapp/tonic/pkg/route"
	"github.com/gin-gonic/gin"
	"github.com/octoper/go-ray"
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
		for k, v := range c.Request.Header {
			fmt.Println("Header: " + k + " " + v[0])
		}
		ray.Ray(c.Request)
		c.JSON(404, gin.H{"code": "P4GE_NOT_FOUND", "message": "Page not found"})
	})
	return r
}