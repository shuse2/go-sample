package services

import (
	"github.com/go-sample/app/models"
)

func GetUser(token string) (models.User, error) {
	// err = errors.New("this is error")
	var err error
	return models.User{}, err
}
