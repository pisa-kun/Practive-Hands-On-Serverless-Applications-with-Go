package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

)

type Idol struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Cover       string `json:"cover"`
	Description string `json:"description"`
}

func findAll() (events.APIGatewayProxyResponse, error) {

	sess := session.New(&aws.Config{Region: aws.String("ap-northeast-1")})
	db := dynamodb.New(sess)
	table_name := "Idols"

	params := &dynamodb.ScanInput{
		TableName: aws.String(table_name), // Required
		AttributesToGet: []*string{
			aws.String("ID"),   // Required
			aws.String("Name"), // Required
			aws.String("Cover"),
			aws.String("Description"),
			// More values...
		},
	}
	res, err := db.Scan(params)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while Scanning DynamoDB",
		}, nil
	}

	idols := make([]Idol, 0)
	for _, item := range res.Items {
		idols = append(idols, Idol{
			ID:          *item["ID"].S,
			Name:        *item["Name"].S,
			Cover:       *item["Cover"].S,
			Description: *item["Description"].S,
		})
	}

	response, err := json.Marshal(idols)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while decoding to string value",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
		Body: string(response),
	}, nil
}

func main() {
	lambda.Start(findAll)
}
