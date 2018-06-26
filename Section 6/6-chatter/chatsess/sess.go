package chatsess

import (
	"crypto/rand"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Login struct {
	Sessid   string
	Username string
}

func NewLogin(name string) Login {
	b := make([]byte, 20)
	rand.Read(b)
	return Login{
		Sessid:   fmt.Sprintf("%x", b),
		Username: name,
	}
}

func GetLogin(id string, sess *session.Session) (Login, error) {

	dbc := dynamodb.New(sess)
	dbres, err := dbc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("ch_sessions"),
		Key:       map[string]*dynamodb.AttributeValue{"Sessid": {S: aws.String(id)}},
	})

	if err != nil {
		return Login{}, err
	}

	if dbres.Item == nil {
		return Login{}, fmt.Errorf("No Item, so null")
	}

	un, ok := dbres.Item["Username"]
	if !ok {
		return Login{}, fmt.Errorf("No Username")
	}

	return Login{Sessid: id, Username: *(un.S)}, nil

}

func (l Login) Put(sess *session.Session) error {
	dbc := dynamodb.New(sess)

	_, err := dbc.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String("ch_sessions"),
		Item: map[string]*dynamodb.AttributeValue{
			"Sessid":   {S: aws.String(l.Sessid)},
			"Username": {S: aws.String(l.Username)},
		},
	})

	return err
}
