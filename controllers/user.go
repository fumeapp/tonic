package controllers

import (
	"net/http"

	"github.com/fumeapp/tonic/models"
	"github.com/fumeapp/tonic/pkg/route"
	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
)

func UserResource() route.ApiResourceStruct {
	return route.ApiResourceStruct{
		Index: func(c *gin.Context) {
			c.JSON(
				http.StatusOK,
				paginate.New().
					With(models.Db.Model(&models.User{}).Preload("Providers")).
					Request(c.Request).Response(&[]models.User{}),
			)
		},
		Show: func(c *gin.Context) {
			var user models.User
			models.Db.Model(&models.User{}).Preload("Providers").First(&user, c.Param("id"))
			c.JSON(http.StatusOK, user)
		},
		Update: func (c *gin.Context) { 
			var user models.User
			models.Db.Model(&models.User{}).Preload("Providers").First(&user, c.Param("id"))
			c.JSON(http.StatusOK, user)
		},
	}
}
