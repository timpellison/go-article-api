package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	message := fmt.Sprintf("Request route: %s, Verb: %s", request.RequestContext.Path, request.RequestContext.HTTPMethod)
	response := fmt.Sprintf("Hello World %s", message)
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: response}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
