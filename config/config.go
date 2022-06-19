package config

import (
	"encoding/json"
	"os"
)

var config Configuration

func LoadConfiguration(filename string) (Configuration, error) {
	configFile, err := os.Open(filename)
	defer configFile.Close()
	if err != nil {
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config, err

}
