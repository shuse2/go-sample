package models

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/golang/glog"
	"gopkg.in/mgo.v2/bson"
)

const (
	TABLE_USER = "user"
)

type User struct {
	UserId   string
	Password string
}

func GetUser(userId string, password string) (*User, error) {
	result := &User{}
	c := dbm.DB(dbNameM).C(TABLE_USER)
	hash := md5.Sum([]byte(password))
	err := c.Find(bson.M{"userid": userId, "password": hex.EncodeToString(hash[:])}).One(result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func CreateUser(user *User) error {
	c := dbm.DB(dbNameM).C(TABLE_USER)
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
