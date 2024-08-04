package config

import (
	"encoding/json"
	"os"
)

const configFile = "config.json"

type Config struct {
	ValidSuffixes []string `json:"suffixes"`
	SearchRoot    string   `json:"searchRoot"`
}

func ParseConfigFile() (*Config, error) {
	rawJson, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	var config Config

	err = json.Unmarshal(rawJson, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
