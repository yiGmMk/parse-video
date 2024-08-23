// netlify函数这里必须是main
package main

import (
	"context"
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	h "github.com/wujunwei928/parse-video/handler"
)

var ginLambda *ginadapter.GinLambda

//go:embed templates/*
var files embed.FS
var router *gin.Engine

type HttpResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func init() {
	router = gin.Default()
	sub, err := fs.Sub(files, "templates")
	if err != nil {
		panic(err)
	}
	h.RegisterHandler(router, sub)

	ginLambda = ginadapter.New(router)
}

// vercel Function
func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	res, err := ginLambda.ProxyWithContext(ctx, req)
	if err != nil {
		log.Default().Println(err)
	}
	return &res, err
}

// netlify
func main() {
	lambda.Start(handler)
}
