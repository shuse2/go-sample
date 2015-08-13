package controllers

import (
	// "fmt"
	"github.com/go-sample/app/models"
	"github.com/golang/glog"
	"github.com/gorilla/sessions"
	"net/http"
)

func TwitterLoginHandler(w http.ResponseWriter, r *http.Request) {
	requestToken, err := models.GetRequestTokenAndURL("http://localhost:3000/twitter/callback")

	if err != nil {
		glog.Info("Failed to obtain request token")
	}

	session, err := models.GetSession(r, "twitter_login")

	if err != nil {
		glog.Error("failed to get session")
	}

	session.Values["request_token"] = requestToken
	if err = sessions.Save(r, w); err != nil {
		glog.Fatalf("Error saving session: %v", err)
	}

	glog.Infof("url: %s", requestToken.Url)
}

func TwitterLoginCallbackHandler(w http.ResponseWriter, r *http.Request) {
	session, err := models.GetSession(r, "twitter_login")
	if err != nil {
		glog.Error("failed to get session")
	}

	val := session.Values["request_token"]

	rToken, ok := val.(*models.RequestToken)
	if !ok {
		glog.Error("failed to parse request token")
	}
	glog.Infof("rtoken: %s", rToken.Token)
	// check req token and session token matches

	aToken, vErr := models.VarifyToken(rToken, r.FormValue("oauth_verifier"))
	if vErr != nil {
		glog.Error("failed to varify request token")
	}

	glog.Infof("access token: %s", aToken.Token)
}
