package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BaseIndex(c *gin.Context, r *gin.Engine) {
	type RouteInfo struct {
		Method string
		Path   string
		Handler  string
	}
	routes := []RouteInfo{}
	for i, e := range r.Routes() {
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
}
