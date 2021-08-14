package main

import (
	"encoding/json"
	"fmt"
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

func insert(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var idol Idol
	err := json.Unmarshal([]byte(request.Body), &idol)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Invalid payload",
		}, nil
	}

	sess := session.New(&aws.Config{Region: aws.String("ap-northeast-1")})
	db := dynamodb.New(sess)
	table_name := "Idols"

	// PutItem
	putParams := &dynamodb.PutItemInput{
		TableName: aws.String(table_name),
		Item: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(idol.ID),
			},
			"Name": {
				S: aws.String(idol.Name),
			},
			"Cover": {
				S: aws.String(idol.Cover),
			},
			"Description": {
				S: aws.String(idol.Description),
			},
		},
	}

	putItem, putErr := db.PutItem(putParams)
	if putErr != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while insert to DynamoDB" + putErr.Error(),
		}, nil
	}
	fmt.Println(putItem)

	response, err := json.Marshal(putItem)
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
	lambda.Start(insert)
}
