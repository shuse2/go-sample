package models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Username string
	Password string
	Token    string
}

func GetUser(token string) (*User, error) {
	result := &User{}
	c := Dbm.DB("go-sample").C("user")
	err := c.Find(bson.M{"token": token}).One(result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func CreateUser(user *User) error {
	c := Dbm.DB("go-sample").C("user")
	return c.Insert(user)
}
