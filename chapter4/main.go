package main

import "github.com/aws/aws-lambda-go/lambda"

type Response struct {
	StatusCode int    `json:"StatusCode"`
	Body       string `json:"Body"`
}

func handler() (Response, error) {
	return Response{
		StatusCode: 728,
		Body:       "Welcome to serverless world",
	}, nil
}

func main() {
	lambda.Start(handler)
}
