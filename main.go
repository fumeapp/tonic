package main

import (
	"net/http"

	"github.com/fumeapp/skele/models"
	"github.com/fumeapp/skele/pkg/setting"
	"github.com/fumeapp/skele/routes"
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
		Addr:    ":8080",
		Handler: routes,
	}

	server.ListenAndServe()


}