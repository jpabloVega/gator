package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadConfig() (Config, error) {
	// Get gator config address
	gatorConfig, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	// Get the json at that address
	config_json, err := os.ReadFile(gatorConfig)
	if err != nil {
		return Config{}, err
	}

	// Print the json
	fmt.Println(string(config_json))

	// Unmarshal json and return it
	config := Config{}
	if err := json.Unmarshal(config_json, &config); err != nil {
		return Config{}, err
	}

	config.Current_user_name = ""

	return config, nil

}
