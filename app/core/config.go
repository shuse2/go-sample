package core

import (
	"encoding/json"
	"io/ioutil"
)

type MongoConfig struct {
	Host     string `json:"host"`
	Database string `json:"database"`
}

type RedisConfig struct {
	Host     string `json:"host"`
	Database int64  `json:"database"`
}

type Configuration struct {
	Secret             string       `json:"secret"`
	PublicPath         string       `json:"public_path"`
	MongoDatabase      MongoConfig  `json:"mongo_config"`
	RedisDatabase      RedisConfig  `json:"redis_config"`
	TwitterLoginConfig TwitterLogin `json:"twitter_login"`
}

type TwitterLogin struct {
	ConsumerKey    string `json:"consumer_key"`
	ConsumerSecret string `json:"consumer_secret"`
}

func (config *Configuration) Load(filename string) error {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	return config.Parse(data)
}

func (config *Configuration) Parse(data []byte) error {
	return json.Unmarshal(data, &config)
}
