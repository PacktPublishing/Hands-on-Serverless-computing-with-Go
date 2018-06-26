package chatsess

import (
	"fmt"
	"html"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type User struct {
	Username string
	Password string
}

func NewUser(name, pass string) User {
	return User{
		Username: html.EscapeString(name),
		Password: NewPassword(pass),
	}
}

func (u User) Put(sess *session.Session) error {
	dbc := dynamodb.New(sess)

	_, err := dbc.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String("ch_users"),
		Item: map[string]*dynamodb.AttributeValue{
			"Username": {S: aws.String(u.Username)},
			"Password": {S: aws.String(u.Password)},
		},
	})

	return err
}

func GetDBUser(uname string, sess *session.Session) (User, error) {
	dbc := dynamodb.New(sess)

	uname = html.EscapeString(uname)

	dbres, err := dbc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("ch_users"),
		Key: map[string]*dynamodb.AttributeValue{
			"Username": {S: aws.String(uname)},
		},
	})

	if err != nil {
		return User{}, err
	}

	if dbres.Item == nil {
		return User{}, fmt.Errorf("No User exists by that name :%s", uname)
	}

	pwv, ok := dbres.Item["Password"]

	if !ok {
		return User{}, fmt.Errorf("User has no password :%s", uname)
	}

	return User{Username: uname, Password: *(pwv.S)}, nil
}

func GetDBUserPass(uname, pass string, sess *session.Session) (User, error) {
	u, err := GetDBUser(uname, sess)
	if err != nil {
		return u, err
	}
	if !CheckPassword(pass, u.Password) {
		return User{}, fmt.Errorf("Password Doesn't Match")
	}

	return u, err
}
