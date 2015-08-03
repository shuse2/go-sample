package models

import (
	"github.com/golang/glog"
	"gopkg.in/mgo.v2"
	"gopkg.in/redis.v3"
)

var (
	dbm     *mgo.Session
	dbr     *redis.Client
	dbNameM string
)

func Init(mongoHost string, mongoDatabase string, redisHost string, redisDatabase int64) {
	mongoC, err := mgo.Dial(mongoHost)

	if err != nil {
		glog.Fatalf("Cannot connect to mongo database %v", err)
		panic(err)
	}
	dbm = mongoC
	dbNameM = mongoDatabase

	redisC := redis.NewClient(&redis.Options{
		Addr: redisHost,
		DB:   redisDatabase,
	})
	dbr = redisC
	// check redis health
	{
		_, err := dbr.Ping().Result()
		if err != nil {
			glog.Fatalf("Cannot connect to redis database %v", err)
			panic(err)
		}
	}
}

func Close() {
	glog.Info("Closing application")
	dbm.Close()
	dbr.Close()
}
