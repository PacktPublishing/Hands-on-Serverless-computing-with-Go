package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/coderconvoy/pakt/chatter/chatsess"
)

type Event struct {
	Sessid string
	Text   string
}

type Response struct {
	Job string
	Err string
}

func handler(c context.Context, ev Event) (Response, error) {
	sess := session.Must(session.NewSession())
	lg, err := chatsess.GetLogin(ev.Sessid, sess)

	if err != nil {
		return Response{Job: "Say " + ev.Text, Err: "Not logged in:" + err.Error()}, nil
	}

	ch := chatsess.NewChat(lg.Username, ev.Text)
	err = ch.Put(sess)
	if err != nil {
		return Response{Job: "Say " + ev.Text, Err: "Could Not:" + err.Error()}, nil
	}

	return Response{Job: "Say " + ev.Text}, nil

}

func main() {
	lambda.Start(handler)
}
