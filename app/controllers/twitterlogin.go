package controllers

import (
	// "fmt"
	"github.com/go-sample/app/models"
	"github.com/golang/glog"
	"net/http"
)

func TwitterLoginHandler(w http.ResponseWriter, r *http.Request) {
	requestToken, err := models.GetRequestTokenAndURL("/twitter/callback")

	if err != nil {
		glog.Info("Failed to obtain request token")
	}

	glog.Infof("url: %s", requestToken.Url)
}
