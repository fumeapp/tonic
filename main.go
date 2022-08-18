package main

import (
	fume "github.com/fumeapp/gin"
	"github.com/fumeapp/tonic/models"
	"github.com/fumeapp/tonic/pkg/setting"
	"github.com/fumeapp/tonic/routes"
)

func init() {

	setting.Setup()
	models.Setup()
	/*
		models.Truncate()
		models.Migrate()
		models.Seed()
	*/
}

func main() {
	routes := routes.Init(setting.IsDev() || setting.IsDebug())
	fume.Start(routes, fume.Options{})
}