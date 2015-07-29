package controllers

import (
	"fmt"
	"github.com/go-sample/app/services"
	"net/http"
)

// controller
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	services.GetUser("aaa")
	fmt.Fprintf(w, "index handler")
}
