package core

import (
	"encoding/json"
	"io/ioutil"
)

type ConfigurationDatabase struct {
	Hosts    string `json:"Hosts"`
	Database string `json:"database"`
}

type Configuration struct {
	Secret     string `json:"seacret"`
	PublicPath string `json:"public_path"`
	Database   ConfigurationDatabase
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
