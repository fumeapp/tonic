package route

import (
	"net/http"
	"reflect"
	"strconv"

	. "github.com/fumeapp/tonic/database"
	"github.com/gin-gonic/gin"
	"github.com/octoper/go-ray"
)

type ApiResourceStruct struct {
	Index  func(c *gin.Context)
	Show   func(c *gin.Context)
	Update func(c *gin.Context)
}

var router *gin.Engine
var modelType reflect.Type
var apirs ApiResourceStruct
var name string

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
	if checkNumeric(c) && retrieveModel(c) {
		ray.Ray("and tell")
		apirs.Show(c)
	}
}

func ApiResource(route *gin.Engine, n string, model any, ctls ApiResourceStruct) {
	apirs = ctls
	modelType = reflect.TypeOf(model)
	name = n
	route.GET("/"+n, ctls.Index)
	route.GET("/"+n+"/:id", show)
	route.PUT("/"+n+"/:id", ctls.Update)
}

func checkNumeric(c *gin.Context) bool {
	if _, err := strconv.Atoi(c.Param("id")); err != nil {
		abortNotFound(c)
		return false
	}
	return true
}

func retrieveModel(c *gin.Context) bool {
	filled := reflect.New(modelType).Interface()
	result := Db.Where("id = ?", c.Param("id")).First(filled)
	if result.Error != nil {
		abortNotFound(c)
		return false
	}
	c.Set(name, filled)
	return true
}

func abortNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"message": "Resource not found"})
}