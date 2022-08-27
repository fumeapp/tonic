package tonic

import (
	"github.com/fumeapp/tonic/aws"
	"github.com/fumeapp/tonic/database"
	"github.com/fumeapp/tonic/route"
	"github.com/fumeapp/tonic/setting"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {

	setting.Setup()
	database.Setup()
	aws.Setup()

	engine := gin.New()
	engine.Use(route.Benchmark)
	return engine
}
