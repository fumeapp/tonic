package route

import (
	"errors"
	"net/http"
	"reflect"
	"strconv"

	. "github.com/fumeapp/tonic/database"
	"github.com/gin-gonic/gin"
)

type ApiResourceStruct struct {
	Index  func(c *gin.Context)
	Show   func(c *gin.Context, filled any)
	Update func(c *gin.Context, filled any)
}

var (
	router    *gin.Engine
	modelType reflect.Type
	apirs     ApiResourceStruct
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

func show(c *gin.Context) {
	if checkNumeric(c) {
		value, error := retrieveModel(c)
		if error != nil {
			abortNotFound(c)
		} else {
			apirs.Show(c, value)
		}
	}
}

func update(c *gin.Context) {
	if checkNumeric(c) {
		value, error := retrieveModel(c)
		if error != nil {
			abortNotFound(c)
		} else {
			apirs.Update(c, value)
		}
	}
}

// note: only pass value to model. dont pass a pointer.
func ApiResource(route *gin.Engine, n string, model any, ctls ApiResourceStruct) {
	apirs = ctls
	modelType = reflect.TypeOf(model)
	route.GET("/"+n, ctls.Index)
	route.GET("/"+n+"/:id", show)
	route.PUT("/"+n+"/:id", update)
}

func checkNumeric(c *gin.Context) bool {
	if _, err := strconv.Atoi(c.Param("id")); err != nil {
		abortNotFound(c)
		return false
	}
	return true
}

func retrieveModel(c *gin.Context) (any, error) {
	model := reflect.New(modelType).Interface()
	result := Db.First(model, c.Param("id"))
	if result.Error != nil {
		abortNotFound(c)
		return -1, errors.New("Record not found")
	}
	return model, nil
}

func abortNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"message": "Resource not found"})
}
