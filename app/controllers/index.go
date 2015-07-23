package controllers

import (
	"fmt"
	"net/http"
)

// controller
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "index handler")
}
