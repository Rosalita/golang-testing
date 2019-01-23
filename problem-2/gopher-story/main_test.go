package main

import (
	"testing"
	"flag"
	"os"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"

)

type mockDynamoDBClient struct {
	dynamodbiface.DynamoDBAPI
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