package models

import (
	"gopkg.in/mgo.v2/bson"
)

type OAuthClient struct {
	Name   string
	Id     string
	Secret string
	UserId string
}

type Token struct {
	UserId   string
	ClientId string
	Token    string
}

const (
	TABLE_OAUTH_CLIENT = "oauth_client"
	TABLE_OAUTH_TOKEN  = "oauth_token"
)

func getOAuthClient(userId string, name string) (*OAuthClient, error) {
	result := &OAuthClient{}
	c := dbm.DB(dbNameM).C(TABLE_OAUTH_CLIENT)
	err := c.Find(bson.M{"name": name, "userid": userId}).One(result)
	if err != nil {
		return result, err
	}
	return result, nil
}
