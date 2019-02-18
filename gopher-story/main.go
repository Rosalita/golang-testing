package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type gopher struct {
	ID     string `dynamodbav:"ID"`
	Name   string `dynamodbav:"Name"`
	Colour string `dynamodbav:"Colour"`
}

var dynamoSession dynamodbiface.DynamoDBAPI

func main() {
	awsConfig := aws.Config{Region: aws.String("eu-west-1")}
	dynamoSession = dynamodb.New(session.New(&awsConfig))

	err := randomStory()

	if err != nil {
		fmt.Println(err)
	}

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
	err = dynamodbattribute.UnmarshalMap(result.Item, &gopher)

	return gopher, nil

}

func tellStory(g gopher) string {
	story := fmt.Sprintf("Once upon a time there was a gopher called %s. Their fur was %s colour", g.Name, g.Colour)
	return story
}

func randomStory() error {

	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(4)
	ID := strconv.Itoa(randNum)

	gopher, err := getGopher(ID)

	if err != nil {
		return err
	}

	log.Println(tellStory(gopher))

	return err

}
