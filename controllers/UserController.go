package UserController

import (
	"net/http"

	"github.com/fumeapp/skele/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var users []models.User
	c.JSON(http.StatusOK, models.Db.Find(&users))
}

func Show(c *gin.Context) {
	var user models.User
	models.Db.First(&user, c.Param("id"))
	c.JSON(http.StatusOK, user)
}