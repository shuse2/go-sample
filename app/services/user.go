package services

import (
	"github.com/go-sample/app/core"
	"github.com/go-sample/app/models"
)

func GetUser(token string) error {
	// err = errors.New("this is error")
	var err error
	user := &models.User{
		Username: "abc",
		Password: "pass",
		Token:    "aaa",
	}
	dbc := core.GetApplicaton()
	c := dbc.DB("go-sample").c("user")
	return c.Insert(user)
}
