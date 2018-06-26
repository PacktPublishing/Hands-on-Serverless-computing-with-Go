package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler(c context.Context, s string) (string, error) {
}

func main() {
	lambda.Start(handler)
}
