package controllers

import (
	"net/http"

	"github.com/fumeapp/tonic/models"
	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
)

func UserIndex(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		paginate.New().With(models.Db.Model(&models.User{})).Request(c.Request).Response(&[]models.User{}),
	)
}

func UserShow(c *gin.Context) {
	var user models.User
 	models.Db.Model(&models.User{}).Preload("Providers").First(&user, c.Param("id"))
	c.JSON(http.StatusOK, user)
}