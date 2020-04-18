package main

import (
	"context"
	"errors"
	"log"

	"github.com/maisiesadler/theilliminationgame"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/maisiesadler/theilliminationgame/apigateway"
)

var errAuth = errors.New("Not logged in")
var errParse = errors.New("Error parsing response")

// GamesResponse is the response from this handler
type GamesResponse struct {
	Games []*theilliminationgame.GameSummary `json:"games"`
}

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)
	log.Printf("Body: '%v'\n", request.Body)

	user, err := apigateway.GetOrCreateAuthenticatedUser(context.TODO(), &request)
	if err != nil {
		return apigateway.ResponseSuccessfulString("Hello unknown"), nil
	}

	resp := apigateway.ResponseSuccessfulString("Hello " + user.Nickname)
	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
