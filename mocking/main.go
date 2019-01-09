package main

import(
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"

)

type gopher struct{
	ID string
	name string
	colour string
}

var dynamoSession dynamodbiface.DynamoDBAPI

func main(){
	awsConfig := aws.Config{Region: aws.String("eu-west-1")}
	dynamoSession = dynamodb.New(session.New(&awsConfig))

	gopher, err := getGopher("1")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(gopher)

}

func getGopher(gopherID string) (gopher, error) {

	getItemInput := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(gopherID),
			},
		},
		TableName: aws.String("gophers"),
	}


	result, err := dynamoSession.GetItem(getItemInput)
	if err != nil {
		return gopher{}, err
	}

	gopher := gopher{}
	if result.Item["ID"] != nil {
		err = dynamodbattribute.UnmarshalMap(result.Item, &gopher)
	}

	return gopher, nil

}
