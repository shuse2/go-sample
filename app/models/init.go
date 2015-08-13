package models

import (
	"github.com/go-sample/app/config"
	"github.com/golang/glog"
	"github.com/gorilla/sessions"
	"gopkg.in/boj/redistore.v1"
	"gopkg.in/mgo.v2"
)

var (
	dbm     *mgo.Session
	store   *redistore.RediStore
	session *sessions.Session
	dbNameM string
)

func Init(conf *config.Configuration) {
	// connection to mongo
	var mgoerr error
	dbm, mgoerr = mgo.Dial(conf.MongoDatabase.Host)

	if mgoerr != nil {
		glog.Fatalf("Cannot connect to mongo database %v", mgoerr)
		panic(mgoerr)
	}
	dbNameM = conf.MongoDatabase.Database

	// session store
	var storeerr error
	store, storeerr = redistore.NewRediStore(10, "tcp", ":6379", "", []byte("secret-key"))
	if storeerr != nil {
		glog.Fatalf("Cannot connect to redis database %v", storeerr)
		panic(storeerr)
	}

	initTwitter(conf)
}

func Close() {
	glog.Info("Closing application")
	dbm.Close()
	store.Close()
}
