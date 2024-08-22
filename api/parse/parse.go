package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	h "github.com/wujunwei928/parse-video/handler"
)

var ginLambda *ginadapter.GinLambda

type HttpResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func init() {
	r := gin.Default()
	h.RegisterHandler(r, "../../templates/*")

	ginLambda = ginadapter.New(r)
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	res, err := ginLambda.ProxyWithContext(ctx, req)
	if err != nil {
		log.Default().Println(err)
	}
	return &res, err
}

func main() {
	lambda.Start(handler)
}
