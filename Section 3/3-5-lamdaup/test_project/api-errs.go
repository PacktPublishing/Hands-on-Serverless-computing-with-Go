package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Username string
}

func handler(e Event) (string, error) {
	if len(e.Username) == 0 {
		return "", fmt.Errorf("No Name Given")
	}

	if e.Username[0] == 'D' {
		return "", fmt.Errorf("Dont Like : %s", e.Username)
	}

	return fmt.Sprintf("<h1>Hello %s from Lambda Go</h1>", e.Username), nil
}

func main() {
	lambda.Start(handler)
}
