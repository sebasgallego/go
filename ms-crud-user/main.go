package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"log"
	"ms-crud-user/database"
	_ "ms-crud-user/docs"
	"ms-crud-user/routes"
	"ms-crud-user/utils"
)

// @title           Epayco PSE API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

var ginLambda *ginadapter.GinLambda

func init() {

	//Init Gin
	r := gin.Default()

	// Connect to database
	database.Connect(utils.GoDotEnvVariable("connection_db"))

	// Migrate database
	database.Migrate()

	//Register routes
	routes.SetupRoutes(r)

	// Run the server local
	//err := r.Run(":" + utils.GoDotEnvVariable("port"))
	err := r.Run()
	if err != nil {
		log.Println("Start sever error...")
		return
	}

	ginLambda = ginadapter.New(r)

}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
