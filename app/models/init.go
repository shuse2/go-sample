package models

import (
	"github.com/go-sample/app/config"
	"github.com/golang/glog"
	"github.com/gorilla/sessions"
	"gopkg.in/boj/redistore.v1"
	"gopkg.in/mgo.v2"
	"net/http"
)

var (
	dbm     *mgo.Session
	store   *redistore.RediStore
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

func GetSession(req *http.Request, key string) (*sessions.Session, error) {
	session, err := store.Get(req, key)
	if err != nil {
		glog.Errorf("Session cannot be obtained %s", err)
		return session, err
	}

	return session, nil
}
