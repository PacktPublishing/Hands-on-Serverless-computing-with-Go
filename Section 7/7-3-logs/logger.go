package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler(c context.Context, s string) (string, error) {
	fmt.Println(s)
	log.Print(s)

	return "Have Logged : " + s, nil
}

func main() {
	lambda.Start(handler)
}
