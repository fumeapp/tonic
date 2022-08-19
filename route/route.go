package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiResourceStruct struct {
	Index  func(c *gin.Context)
	Show   func(c *gin.Context)
	Update func(c *gin.Context)
}

var router *gin.Engine

func Routes(route *gin.Engine) {
	router = route
	route.GET("/", RouteList)
}

func RouteList(c *gin.Context) {
	type RouteInfo struct {
		Method  string
		Path    string
		Handler string
	}
	routes := []RouteInfo{}
	for _, e := range router.Routes() {
		routes = append(routes, RouteInfo{
			Method:  e.Method,
			Path:    e.Path,
			Handler: e.Handler,
		})
	}
	c.JSON(
		http.StatusOK,
		routes,
	)
}

func ApiResource(route *gin.Engine, resource string, ctls ApiResourceStruct) {
	route.GET("/"+resource, ctls.Index)
	route.GET("/"+resource+"/:id", ctls.Show)
	route.PUT("/"+resource+"/:id", ctls.Update)
}
