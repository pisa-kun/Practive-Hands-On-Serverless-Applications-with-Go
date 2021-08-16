package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

func Cloudwatch() {
	// insert logic

	sess := session.New(&aws.Config{Region: aws.String("ap-northeast-1")})
	svc := cloudwatch.New(sess)

	_, err := svc.PutMetricData(&cloudwatch.PutMetricDataInput{
		Namespace: aws.String("InsertIdol"),
		MetricData: []*cloudwatch.MetricDatum{
			&cloudwatch.MetricDatum{
				MetricName: aws.String("ActionIdol"),
				Unit:       aws.String("Count"),
				Value:      aws.Float64(5885.0),
				Dimensions: []*cloudwatch.Dimension{
					&cloudwatch.Dimension{
						Name:  aws.String("Environment"),
						Value: aws.String("staging"),
					},
				},
			},
		},
	})
	if err != nil {
		fmt.Println("Error adding metrics:", err.Error())
		return
	}
}

func main() {
	lambda.Start(Cloudwatch)
}
