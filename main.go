package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type testEvent struct {
	Name string `json:"name"`
}

func LambdaHandler(name testEvent) (string, error) {
	return fmt.Sprintf("Hello %s!", name.Name), nil
}

func main() {
	lambda.Start(LambdaHandler)
}
