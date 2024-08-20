package main

import (
	"github.com/JoseLuis21/skeleton-lambda-ms/internal/lambda_handler"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	handlerLambda := lambda_handler.NewHandler()
	lambda.Start(handlerLambda.HandlerRequest)
}
