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
	ID   string `json:"id"`
	Name string `json:"name"`
}

func deleteItem(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var idol Idol
	err := json.Unmarshal([]byte(request.Body), &idol)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       err.Error(),
		}, nil
	}

	sess := session.New(&aws.Config{Region: aws.String("ap-northeast-1")})
	db := dynamodb.New(sess)
	table_name := "Idols"

	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(table_name),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(idol.ID),
			},
		},
	}
	// Execute.
	res, err := db.DeleteItem(input)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while delete item to DynamoDB" + err.Error(),
		}, nil
	}

	response, err := json.Marshal(res)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while decoding to string value",
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
	lambda.Start(deleteItem)
}
