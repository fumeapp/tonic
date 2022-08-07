package usercontroller

import (
	"net/http"

	"github.com/fumeapp/skele/models"
	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
)

func Index(c *gin.Context) {
	pg := paginate.New()
	model := models.Db.Model(&models.User{})
	page := pg.Response(model, c.Request, &[]models.User{})
	c.JSON(http.StatusOK, page)
}

func Show(c *gin.Context) {
	var user models.User
	models.Db.First(&user, c.Param("id"))
	c.JSON(http.StatusOK, user)
}