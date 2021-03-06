package main

import (
	"flag"
	"os"
	"testing"
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/stretchr/testify/assert"
)

type mockDynamoDBClient struct {
	dynamodbiface.DynamoDBAPI
}

func (m *mockDynamoDBClient) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	var result dynamodb.GetItemOutput

	switch *input.TableName {
	case "gophers":
		if input.Key["ID"].S != nil && *input.Key["ID"].S == "1"{
			result = dynamodb.GetItemOutput{
				Item: map[string]*dynamodb.AttributeValue{
					"ID":     {S: aws.String("1")},
					"Name":   {S: aws.String("Bob")},
					"Colour": {S: aws.String("Brown")},
				},
			}
		}
		if input.Key["ID"].S != nil && *input.Key["ID"].S == "err"{
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
	dynamoSession = &mockDynamoDBClient{}

	// Run tests
	exitCode := m.Run()

	// Test teardown
	os.Exit(exitCode)
}

func TestGetGopher(t *testing.T) {

	expectedGopher := gopher{
		ID:     "1",
		Name:   "Bob",
		Colour: "Brown",
	}

	tests := []struct {
		gopherID string
		expected gopher
		err      error
	}{
		{"1", expectedGopher, nil},
		{"err", gopher{}, errors.New("Dynamo Internal Server Error")},
	}

	for _, test := range tests {
		result, err := getGopher(test.gopherID)
		assert.Equal(t, test.expected, result)
		assert.Equal(t, test.err, err)
	}
}
