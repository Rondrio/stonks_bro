package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Token string `json:"token"`
}

func ReadConfig(path string) (*Config, error) {
	c, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg := Config{}

	err = json.Unmarshal(c, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
