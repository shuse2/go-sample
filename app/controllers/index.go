package controllers

import (
	"fmt"
	"github.com/go-sample/app/models"
	"net/http"
)

// controller
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	user := &models.User{
		UserId:   "aaa",
		Password: "12345",
	}
	models.CreateUser(user)
	fmt.Fprintf(w, "index handler")
}
