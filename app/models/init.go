package models

import (
	"github.com/golang/glog"
	"gopkg.in/mgo.v2"
)

var (
	dbm      *mgo.Session
	database string
)

func Init(host string, dbName string) {
	connection, err := mgo.Dial(host)

	if err != nil {
		glog.Fatalf("Cannot connect to database %v", err)
		panic(err)
	}
	dbm = connection
	database = dbName
}

func Close() {
	glog.Info("Closing application")
	dbm.Close()
}
