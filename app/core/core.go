package core

import (
	"github.com/golang/glog"
	"github.com/gorilla/sessions"
	"gopkg.in/mgo.v2"
)

type Application struct {
	Configuration *Configuration
	Store         *sessions.CookieStore
	DBSession     *mgo.Session
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

func (application *Application) ConnectDB() {
	var err error
	application.DBSession, err = mgo.Dial(application.Configuration.Database.Hosts)

	if err != nil {
		glog.Fatalf("Cannot connect to database %v", err)
		panic(err)
	}
}

func (application *Application) Close() {
	glog.Info("Closing application")
	application.DBSession.Close()
}
