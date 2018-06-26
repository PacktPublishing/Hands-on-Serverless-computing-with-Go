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
	Job    string
	Sessid string
	Err    string
}

func handler(c context.Context, ev Event) (Response, error) {
	sess := session.Must(session.NewSession())

	u, err := chatsess.GetDBUserPass(ev.Username, ev.Password, sess)

	if err != nil {
		return Response{
			Job: "Login", Err: err.Error(),
		}, nil
	}

	lg := chatsess.NewLogin(u.Username)
	err = lg.Put(sess)

	if err != nil {
		return Response{Job: "Login", Err: err.Error()}, nil
	}

	return Response{Job: "Login", Sessid: lg.Sessid}, nil
}

func main() {
	lambda.Start(handler)
}
