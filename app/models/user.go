package models

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/golang/glog"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	UserId   string
	Password string
}

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

func GetUser(userId string, password string) (*User, error) {
	result := &User{}
	c := dbm.DB(database).C("user")
	hash := md5.Sum([]byte(password))
	err := c.Find(bson.M{"userid": userId, "password": hex.EncodeToString(hash[:])}).One(result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func GetOAuthClient(userId string) ([]OAuthClient, error) {
	var results []OAuthClient
	c := dbm.DB(database).C("oAuthClient")
	err := c.Find(bson.M{"userId": userId}).All(results)

	if err != nil {
		return results, err
	}
	return results, nil
}

func CreateUser(user *User) error {
	c := dbm.DB(database).C("user")
	hashPassword(user)
	res := c.Insert(user)
	if res != nil {
		glog.Info(res.Error())
	}
	return c.Insert(user)
}

func hashPassword(user *User) {
	hash := md5.Sum([]byte(user.Password))
	user.Password = hex.EncodeToString(hash[:])
}
