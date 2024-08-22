package parse

import (
	"context"
	"embed"
	"io/fs"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	h "github.com/wujunwei928/parse-video/handler"
)

var (
	ginLambda *ginadapter.GinLambda

	//go:embed templates/*
	files embed.FS
)

type HttpResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func init() {
	r := gin.Default()
	sub, err := fs.Sub(files, "templates")
	if err != nil {
		panic(err)
	}
	h.RegisterHandler(r, sub)

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
