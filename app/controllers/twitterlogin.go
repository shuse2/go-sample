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
		http.Error(w, "Failed to obtain request token", http.StatusInternalServerError)
	}

	session, err := models.GetSession(r, "twitter_login")

	if err != nil {
		glog.Error("failed to get session")
		http.Error(w, "failed to get session", http.StatusInternalServerError)
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
		http.Error(w, "failed to parse request token", http.StatusInternalServerError)
	}

	token := r.FormValue("oauth_token")
	if token != rToken.Token {
		glog.Error("request token not match")
		http.Error(w, "token not match", http.StatusInternalServerError)
	}
	glog.Infof("rtoken: %s", rToken.Token)

	aToken, vErr := models.VarifyToken(rToken, r.FormValue("oauth_verifier"))
	if vErr != nil {
		glog.Error("failed to varify request token")
		http.Error(w, "failed to varify request token", http.StatusMethodNotAllowed)
	}

	glog.Infof("access token: %s", aToken.Token)
}
