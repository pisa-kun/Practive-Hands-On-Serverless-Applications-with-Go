package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-xray-sdk-go/xray"
)

type Idol struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Cover       string `json:"cover"`
	Description string `json:"description"`
}

func findAll(ctx context.Context) (events.APIGatewayProxyResponse, error) {
	xray.Configure(xray.Config{
		LogLevel:       "info",
		ServiceVersion: "1.2.3",
	})

	sess := session.New(&aws.Config{Region: aws.String("ap-northeast-1")})
	db := dynamodb.New(sess)
	xray.AWS(db.Client)
	table_name := "Idols"

	res, err := db.ScanWithContext(ctx, &dynamodb.ScanInput{
		TableName: aws.String(table_name),
	})

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while scanning DynamoDB",
		}, err
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
