package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"

	"github.com/fumeapp/tonic/models"
	"github.com/fumeapp/tonic/pkg/setting"
	"github.com/fumeapp/tonic/routes"
)

var ginLambda *ginadapter.GinLambda

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

	routes := routes.Init(setting.IsDev())

	if (setting.IsDev()) {
		server := &http.Server{
			Addr:    ":8080",
			Handler: routes,
		}
		server.ListenAndServe()
	} else {
		ginLambda = ginadapter.New(routes)
		lambda.Start(Handler)
	}

}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}