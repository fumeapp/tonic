package route

import (
	"net/http"
	"strconv"

	"errors"

	"github.com/fumeapp/tonic/database"
	"github.com/gin-gonic/gin"
)

type ApiResourceStruct struct {
	Index  func(c *gin.Context)
	Show   func(c *gin.Context, value any)
	Update func(c *gin.Context, value any)
}

var (
	router *gin.Engine
	model any
	resources ApiResourceStruct
)

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
	c.JSON(http.StatusOK, routes)
}

func show (c *gin.Context) {
	if isNumeric(c) {
		value, error := retrieve(c)
		if error != nil {
			invalid(c)
		} else {
			resources.Show(c, value)
		}
	} else {
		invalid(c)
	}
}

func update (c *gin.Context) {
	if isNumeric(c)  {
		value, error := retrieve(c)
		if error != nil {
			invalid(c)
		} else {
			resources.Update(c, value)
		}
	} else {
		invalid(c)
	}
}

func ApiResource(route *gin.Engine, n string, _model any, _resources ApiResourceStruct) {
	resources = _resources
	model = _model
	route.GET("/"+n, resources.Index)
	route.GET("/"+n+"/:id", show)
	route.PUT("/"+n+"/:id", update)
}

func isNumeric(c *gin.Context) bool {
	if _, err := strconv.Atoi(c.Param("id")); err != nil {
		return false
	}
	return true
}

func retrieve(c *gin.Context) (any, error) {
	result := database.Db.First(&model, c.Param("id"))
	if result.Error != nil {
		return -1, errors.New("Record not found")
	}
	return model, nil
}

func invalid(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"message": "Resource not found"})
}