package main

import (
	"net/http"

	"github.com/fumeapp/tonic/models"
	"github.com/fumeapp/tonic/pkg/setting"
	"github.com/fumeapp/tonic/routes"
)


func init() {
	setting.Setup()
	models.Setup()
	models.Truncate()
	models.Migrate()
	models.Seed()
}

func main() {

	routes := routes.Init()

	server := &http.Server{
		Addr:    ":8000",
		Handler: routes,
	}

	server.ListenAndServe()


}