package main

import (
	"errors"
	"flag"
	"os"
	"testing"
	"bytes"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/stretchr/testify/assert"
)

func getFakeRandNum(x int) int {
	return 1
}

type mockDynamoDBClient struct {
	dynamodbiface.DynamoDBAPI
}

func (m *mockDynamoDBClient) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	var result dynamodb.GetItemOutput

	switch *input.TableName {
	case "gophers":
		if input.Key["ID"].S != nil && *input.Key["ID"].S == "1" {
			result = dynamodb.GetItemOutput{
				Item: map[string]*dynamodb.AttributeValue{
					"ID":     {S: aws.String("1")},
					"Name":   {S: aws.String("Bob")},
					"Colour": {S: aws.String("Brown")},
				},
			}
		}
		if input.Key["ID"].S != nil && *input.Key["ID"].S == "err" {
			testErr := errors.New("Dynamo Internal Server Error")
			return &dynamodb.GetItemOutput{}, testErr

		}

	default:
		result = dynamodb.GetItemOutput{}
	}

	return &result, nil
}

func TestMain(m *testing.M) {

	flag.Parse()
	// Test setup
	getRandomNum = getFakeRandNum
	dynamoSession = &mockDynamoDBClient{}

	// Run tests
	exitCode := m.Run()

	// Test teardown
	os.Exit(exitCode)
}

func TestTellStory(t *testing.T) {
	emptyGopher := gopher{}
	emptyGopherExpected := "Once upon a time there was a gopher called . Their fur was  colour"

	gopherWithName := gopher{
		Name: "Rosie",
	}
	gopherWithNameExpected := "Once upon a time there was a gopher called Rosie. Their fur was  colour"

	gopherWithColour := gopher{
		Colour: "Red",
	}
	gopherWithColourExpected := "Once upon a time there was a gopher called . Their fur was Red colour"

	gopherWithNameAndColour := gopher{
		Name:   "Rosalita",
		Colour: "Pink",
	}
	gopherWithNameAndColourExpected := "Once upon a time there was a gopher called Rosalita. Their fur was Pink colour"

	tests := []struct {
		input  gopher
		result string
	}{
		{emptyGopher, emptyGopherExpected},
		{gopherWithName, gopherWithNameExpected},
		{gopherWithColour, gopherWithColourExpected},
		{gopherWithNameAndColour, gopherWithNameAndColourExpected},
	}

	for _, test := range tests {
		result := tellStory(test.input)
		assert.Equal(t, test.result, result)
	}

}

func TestRandomStory(t *testing.T) {

	var buffer bytes.Buffer
	log.SetOutput(&buffer)

	tests := []struct {
	
		logged string
		err error
	}{
		{"Once upon a time there was a gopher called Bob. Their fur was Brown colour",nil},
	}

	for _, test := range tests {
		err := randomStory()
		logged := buffer.String()
		assert.Contains(t, logged, test.logged)
		assert.Equal(t, test.err, err)
		buffer.Reset()
	}
}
