package main

import (
	"context"
	"crypto/rand"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Cookie string
}

type Response struct {
	Cookie string
	Body   string
}

func handler(c context.Context, ev Event) (Response, error) {
	cb := make([]byte, 10)
	rand.Read(cb)
	cookie := fmt.Sprintf("%x", cb)
	body := "<h1>Site</h1>\n"
	body += "<p>Old Cookie:" + ev.Cookie + "</p>\n"
	body += "<p>New Cookie:" + cookie + "</p>\n"

	return Response{
		Body:   body,
		Cookie: "sessid=" + cookie,
	}, nil
}

func main() {
	lambda.Start(handler)
}
