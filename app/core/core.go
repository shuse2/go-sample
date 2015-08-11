package core

import (
	"github.com/go-sample/app/config"
	"github.com/golang/glog"
	"github.com/gorilla/sessions"
)

type Application struct {
	Configuration *config.Configuration
	Store         *sessions.CookieStore
}

func (application *Application) Init(filename string) {
	application.Configuration = &config.Configuration{}

	err := application.Configuration.Load(filename)

	if err != nil {
		glog.Fatalf("Cannot read configuration file %s", filename)
		panic(err)
	}

	application.Store = sessions.NewCookieStore([]byte(application.Configuration.Secret))
}
