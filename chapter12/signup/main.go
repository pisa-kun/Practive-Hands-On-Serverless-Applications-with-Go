package main

import (
	"log"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/aws"
)

// ユーザーアカウントを追加するスクリプト
func main() {
	sess := session.New(&aws.Config{Region: aws.String("ap-northeast-1")})

	cognito := cognitoidentityprovider.New(sess)
	req, resp := cognito.SignUpRequest(&cognitoidentityprovider.SignUpInput{
		// 全般設定 -> アプリクライアント
		ClientId: aws.String(os.Getenv("COGNITO_CLIENT_ID")),
		Username: aws.String("EMAIL"),
		Password: aws.String("PASSWORD"),
	})

	fmt.Println(resp)
	err := req.Send()
	if err != nil {
		log.Fatal(err)
	}else{
		fmt.Println("create user account")
	}
}