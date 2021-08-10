package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

)

type Idol struct {
	ID   int    `json:"ID, string"`
	Name string `json:"name"`
}

func main() {

	sess := session.New(&aws.Config{Region: aws.String("ap-northeast-1")})
	db := dynamodb.New(sess)
	table_name := "Idols"

	idols, err := readIdols("idols.json")
	if err != nil {
		log.Fatal(err)
	}

	for _, idol := range idols {
		fmt.Println("Inserting:", idol.Name)
		params := &dynamodb.PutItemInput{
			TableName: aws.String(table_name),
			Item: map[string]*dynamodb.AttributeValue{
				"ID": {
					S: aws.String(strconv.Itoa(idol.ID)),
				},
				"Name": {
					S: aws.String(idol.Name),
				},
			},
		}

		_, err := db.PutItem(params)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func readIdols(fileName string) ([]Idol, error) {
	idols := make([]Idol, 0)

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return idols, err
	}

	fmt.Println(string(data))
	err = json.Unmarshal(data, &idols)
	if err != nil {
		return idols, err
	}

	return idols, nil
}
