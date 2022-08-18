package controllers

import (
	"net/http"

	"github.com/fumeapp/tonic/models"
	"github.com/fumeapp/tonic/pkg/route"
	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
)

func index(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		paginate.New().
			With(models.Db.Model(&models.User{}).Preload("Providers")).
			Request(c.Request).Response(&[]models.User{}),
	)
}

func show(c *gin.Context) {
	var user models.User
	models.Db.Model(&models.User{}).Preload("Providers").First(&user, c.Param("id"))
	c.JSON(http.StatusOK, user)
}

func update(c *gin.Context) {
	var user models.User
	models.Db.Model(&models.User{}).Preload("Providers").First(&user, c.Param("id"))
	c.JSON(http.StatusOK, user)
}

func UserResource() route.ApiResourceStruct {
	return route.ApiResourceStruct{Index: index, Show: show, Update: update}
}
