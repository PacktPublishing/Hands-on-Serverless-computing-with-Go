package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/coderconvoy/pakt/chatter/chatsess"
)

type Event struct {
	Username string
	Password string
}

type Response struct {
	Job string
	Err string
}

func handler(c context.Context, ev Event) (Response, error) {
	sess := session.Must(session.NewSession())

	_, err := chatsess.GetDBUser(ev.Username, sess)

	if err == nil {
		return Response{Job: "Add User", Err: "User already exists"}, nil
	}

	u := chatsess.NewUser(ev.Username, ev.Password)
	err = u.Put(sess)

	if err != nil {
		return Response{Job: "Add User", Err: "Could not add to DB : " + err.Error()}, nil
	}

	return Response{Job: "Add User"}, nil
}

func main() {
	lambda.Start(handler)
}
