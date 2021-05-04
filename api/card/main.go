package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rafaelwi/github-readme-lastfm-stats/src/fetcher"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Get user's name
	res := "Hello github-readme-lastfm-stats "
	if user, userOk := request.QueryStringParameters["user"]; userOk {
		res += "User " + user
	} else {
		res += "No user given"
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "image/svg+xml"},
		Body:       fetcher.SendTestResponse(),
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
