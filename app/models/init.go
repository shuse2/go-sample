package models

import (
	"github.com/go-sample/app/config"
	"github.com/golang/glog"
	"gopkg.in/mgo.v2"
	"gopkg.in/redis.v3"
)

var (
	dbm     *mgo.Session
	dbr     *redis.Client
	dbNameM string
)

func Init(conf *config.Configuration) {
	mongoC, err := mgo.Dial(conf.MongoDatabase.Host)

	if err != nil {
		glog.Fatalf("Cannot connect to mongo database %v", err)
		panic(err)
	}
	dbm = mongoC
	dbNameM = conf.MongoDatabase.Database

	redisC := redis.NewClient(&redis.Options{
		Addr: conf.RedisDatabase.Host,
		DB:   conf.RedisDatabase.Database,
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

	initTwitter(conf)
}

func Close() {
	glog.Info("Closing application")
	dbm.Close()
	dbr.Close()
}
