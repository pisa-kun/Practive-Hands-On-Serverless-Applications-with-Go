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

func insert(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var idol Idol
	//body := []byte(`{"id":"8","name":"tukioka kogane"}`)
	//err := json.Unmarshal(body, &idol)
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

	// Create an input.
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(table_name),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(idol.ID),
			},
			// "Name": {
			// 	S: aws.String(idol.Name),
			// },
		},
		ExpressionAttributeNames: map[string]*string{
			"#Name": aws.String("Name"), // 項目名をプレースホルダに入れる
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":name_value": {
				S: aws.String(idol.Name), // 値をプレースホルダに入れる
			},
		},
		UpdateExpression: aws.String("set #Name = :name_value"), //プレースホルダを利用して更新の式を書く

		//あとは返してくる情報の種類を指定する
		ReturnConsumedCapacity:      aws.String("NONE"),
		ReturnItemCollectionMetrics: aws.String("NONE"),
		ReturnValues:                aws.String("NONE"),
	}
	// Execute.
	res, err := db.UpdateItem(input)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while insert to DynamoDB" + err.Error(),
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
	lambda.Start(insert)
}
