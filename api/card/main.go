package main

import (
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rafaelwi/github-readme-lastfm-stats/src/fetcher"
	"github.com/rafaelwi/github-readme-lastfm-stats/src/generator"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var res string
	showScrobbles := false

	// Get username, get last.fm data, generate card.
	// Return nothing at any time if a step fails.
	user, userOk := request.QueryStringParameters["user"]
	theme, themeOk := request.QueryStringParameters["theme"]
	showScrobblesParam, ssOk := request.QueryStringParameters["show_scrobbles"]

	if ssOk {
		showScrobbles, _ = strconv.ParseBool(showScrobblesParam)
	}

	if !themeOk {
		theme = "light"
	}

	if userOk {
		lastfmData, err := fetcher.GetLastfmData(user, os.Getenv("LASTFM_STATS_KEY"))
		if err != nil {
			res = ""
		}
		res = generator.GenerateCard(lastfmData, theme, showScrobbles)
	} else {
		res = ""
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "image/svg+xml"},
		Body:       res,
	}, nil
}

func main() {
	lambda.Start(handler)
}
