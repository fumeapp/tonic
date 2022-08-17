package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiResourceStruct struct {
	Index	func(c *gin.Context)
	Show	func(c *gin.Context)
	Update	func(c *gin.Context)
}

func Routes(route *gin.Engine) {
	route.GET("/", func(c *gin.Context) {
		type RouteInfo struct {
			Method string
			Path   string
			Handler  string
		}
		routes := []RouteInfo{}
		for _, e := range route.Routes() {
			routes = append(routes, RouteInfo{
				Method: e.Method,
				Path:   e.Path,
				Handler: e.Handler,
			})
		}
		c.JSON(
			http.StatusOK,
			routes,
		)
	})
}

func ApiResource(route *gin.Engine, resource string, ctls ApiResourceStruct) {
	route.GET("/" + resource, ctls.Index)
	route.GET("/" + resource + "/:id", ctls.Show)
	route.PUT("/" + resource + "/:id", ctls.Update)
}