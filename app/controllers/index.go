package controllers

import (
	"fmt"
	"github.com/go-sample/app/models"
	"net/http"
)

// controller
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	models.GetUser("aaa")
	fmt.Fprintf(w, "index handler")
}
