package models

import (
	"github.com/go-sample/app/config"
	"github.com/golang/glog"
	"github.com/mrjones/oauth"
)

var (
	consumer *oauth.Consumer
)

func initTwitter(conf *config.Configuration) {
	consumer = oauth.NewConsumer(
		conf.TwitterLoginConfig.ConsumerKey, conf.TwitterLoginConfig.ConsumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		})
}

func GetRequestTokenAndURL(userId string) (string, error) {
	requestToken, url, err := consumer.GetRequestTokenAndUrl("oob")
	if err != nil {
		glog.Fatal("Twitter Failed to get request token")
		return "", err
	}
	client := &OAuthClient{
		Name:   "Twitter",
		Id:     requestToken.Token,
		Secret: requestToken.Secret,
		UserId: userId,
	}

	cm := dbm.DB(dbNameM).C(TABLE_OAUTH_CLIENT)
	dberr := cm.Insert(client)

	if dberr != nil {
		return "", dberr
	}
	return url, nil
}

func VarifyToken(userId string, varificationCode string) error {

	oauthclient, err := getOAuthClient(userId, "Twitter")

	if err != nil {
		glog.Error("oauthclient not found")
		return err
	}

	requestToken := &oauth.RequestToken{
		Token:  oauthclient.Id,
		Secret: oauthclient.Secret,
	}

	accessToken, err := consumer.AuthorizeToken(requestToken, varificationCode)
	if err != nil {
		return err
	}

	token := &Token{
		UserId:   userId,
		ClientId: oauthclient.Id,
		Token:    accessToken.Token,
	}

	cm := dbm.DB(dbNameM).C(TABLE_OAUTH_TOKEN)
	dberr := cm.Insert(token)

	if dberr != nil {
		glog.Errorf("oauthclient not found: %s", dberr.Error)
		return dberr
	}

	return nil
}
