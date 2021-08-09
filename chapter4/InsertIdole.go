package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Idole struct {
	ID   int    `json:"id`
	Name string `json:"name"`
}

var houkura = []Idole{
	Idole{
		ID:   1,
		Name: "Komiya Kaho",
	},
	Idole{
		ID:   2,
		Name: "Sonoda Chiyoko",
	},
	Idole{
		ID:   3,
		Name: "Saijo jyuri",
	},
	Idole{
		ID:   4,
		Name: "Morino Rinze",
	},
	Idole{
		ID:   5,
		Name: "Arisugawa Natsuha",
	},
}

func insert(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var idole Idole
	err := json.Unmarshal([]byte(req.Body), &idole)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Invalid payload",
		}, nil
	}

	houkura = append(houkura, idole)

	response, err := json.Marshal(houkura)
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
	lambda.Start(insert)
}
