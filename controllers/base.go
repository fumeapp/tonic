package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ControllerInterface interface {
	Index()
	Show()
	Update()
	Delete()
}

func BaseIndex(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		"base controller",
	)
}
