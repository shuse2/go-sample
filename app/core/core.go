package core

import (
	"github.com/golang/glog"
	"github.com/gorilla/sessions"
)

type Application struct {
	Configuration *Configuration
	Store         *sessions.CookieStore
}

func (application *Application) Init(filename string) {
	application.Configuration = &Configuration{}

	err := application.Configuration.Load(filename)

	if err != nil {
		glog.Fatalf("Cannot read configuration file %s", filename)
		panic(err)
	}

	application.Store = sessions.NewCookieStore([]byte(application.Configuration.Secret))
}
