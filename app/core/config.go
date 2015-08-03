package core

import (
	"encoding/json"
	"io/ioutil"
)

type ConfigurationDatabase struct {
	Host     string `json:"host"`
	Database string `json:"database"`
}

type Configuration struct {
	Secret     string                `json:"secret"`
	PublicPath string                `json:"public_path"`
	Database   ConfigurationDatabase `json:"configuration_database"`
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
