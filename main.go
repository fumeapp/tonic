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

var ginLambda *ginadapter.GinLambdaV2

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

	if (setting.IsDev()) {
		server := &http.Server{
			Addr:    ":8000",
			Handler: routes,
		}
		server.ListenAndServe()
	} else {
		ginLambda = ginadapter.NewV2(routes)
		lambda.Start(Handler)
	}

}

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}