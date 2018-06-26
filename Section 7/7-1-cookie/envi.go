package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler(c context.Context) (string, error) {
	env := os.Getenv("MyEnv")
	return "Hello to " + env, nil
}

func main() {
	lambda.Start(handler)
}
