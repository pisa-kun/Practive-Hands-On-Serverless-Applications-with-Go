package main

import (
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var houkura = []struct {
	ID   int    `json:"id`
	Name string `json:"name"`
}{
	{
		ID:   1,
		Name: "Komiya Kaho",
	},
	{
		ID:   2,
		Name: "Sonoda Chiyoko",
	},
	{
		ID:   3,
		Name: "Saijo jyuri",
	},
	{
		ID:   4,
		Name: "Morino Rinze",
	},
	{
		ID:   5,
		Name: "Arisugawa Natsuha",
	},
}

func findOne(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id, err := strconv.Atoi(req.PathParameters["id"])
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "ID Must be a number",
		}, nil
	}

	response, err := json.Marshal(houkura[id-1])
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(response),
	}, nil
}

func main() {
	lambda.Start(findOne)
}
