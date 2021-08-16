package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

// 呼び出し元のLambdaから受け取るための構造体定義
type Event struct {
	Name string `json:"name"`
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func handler(event Event) (string, error) {
	log.Println("Before:", event.Name)
	output := reverse(event.Name)
	log.Println("After:", output)
	return output, nil
}

func main() {
	lambda.Start(handler)
}
