package main

import (
	"context"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/coderconvoy/pakt/chatter/chatsess"
)

type Event struct {
	Sessid   string
	LastID   string
	LastTime string
}

type Response struct {
	Job   string
	Err   string
	Chats []chatsess.Chat
}

func handler(c context.Context, ev Event) (Response, error) {
	sess := session.Must(session.NewSession())

	_, err := chatsess.GetLogin(ev.Sessid, sess)
	if err != nil {
		return Response{Job: "Read", Err: "Not Logged in:" + err.Error()}, nil
	}

	if ev.LastID != "" {
		ltime, err := time.Parse(time.RFC3339, ev.LastTime)
		if err != nil {
			return Response{Job: "ReadL", Err: err.Error()}, nil
		}

		ch, err := chatsess.GetChatAfter(ev.LastID, ltime, sess)
		if err != nil {
			return Response{Job: "ReadL", Err: err.Error()}, nil
		}
		return Response{Job: "ReadL", Chats: ch}, nil
	}

	ch, err := chatsess.GetChat(sess)
	if err != nil {
		return Response{Job: "Read", Err: err.Error()}, nil
	}
	return Response{Job: "Read", Chats: ch}, nil

}

func main() {
	lambda.Start(handler)
}
