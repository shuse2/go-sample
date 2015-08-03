package controllers

import (
	// "fmt"
	"github.com/go-sample/app/models"
	"github.com/golang/glog"
	"html/template"
	"net/http"
)

// controller
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "index handler")
	index := template.Must(template.ParseFiles("views/index.html"))
	index.ExecuteTemplate(w, "index.html", nil)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	user := &models.User{
		UserId:   "aaa",
		Password: "12345",
	}
	models.CreateUser(user)
	glog.Info("user called")
}
