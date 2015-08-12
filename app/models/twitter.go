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

func GetRequestTokenAndURL(callbackUrl string) (*RequestToken, error) {
	requestToken, url, err := consumer.GetRequestTokenAndUrl(callbackUrl)
	rToken := &RequestToken{}
	if err != nil {
		glog.Fatal("Twitter Failed to get request token")
		return rToken, err
	}

	rToken.Token = requestToken.Token
	rToken.Secret = requestToken.Secret
	rToken.Url = url

	return rToken, nil
}

func VarifyToken(requestToken *RequestToken, varificationCode string) (*AccessToken, error) {

	aToken := &AccessToken{}

	rToken := &oauth.RequestToken{
		Token:  requestToken.Token,
		Secret: requestToken.Secret,
	}

	accessToken, err := consumer.AuthorizeToken(rToken, varificationCode)
	if err != nil {
		return aToken, err
	}
	aToken.Token = accessToken.Token
	aToken.Secret = accessToken.Secret
	accessToken.AdditionalData = accessToken.AdditionalData

	return aToken, nil
}
