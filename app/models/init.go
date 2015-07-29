package models

import (
	"github.com/golang/glog"
	"gopkg.in/mgo.v2"
)

var (
	Dbm *mgo.Session
)

func Init(filename string) {
	dbm, err := mgo.Dial(application.Configuration.Database.Hosts)

	if err != nil {
		glog.Fatalf("Cannot connect to database %v", err)
		panic(err)
	}
}

func Close() {
	glog.Info("Closing application")
	application.DBSession.Close()
}
