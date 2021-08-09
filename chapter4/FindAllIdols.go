package main

import (
	"encoding/json"

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

func findAll() (events.APIGatewayProxyResponse, error) {
	response, err := json.Marshal(houkura)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
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
	lambda.Start(findAll)
}
