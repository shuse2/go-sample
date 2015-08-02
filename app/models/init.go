package models

import (
	"github.com/golang/glog"
	"gopkg.in/mgo.v2"
)

var (
	Dbm *mgo.Session
)

func Init(host string) {
	dbm, err := mgo.Dial(host)

	if err != nil {
		glog.Fatalf("Cannot connect to database %v", err)
		panic(err)
	}
	Dbm = dbm
}

func Close() {
	glog.Info("Closing application")
	Dbm.Close()
}
