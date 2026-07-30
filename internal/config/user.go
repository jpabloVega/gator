package config

import (
	"encoding/json"
	"os"
)

func (c *Config) SetUser(name string) error {
	// Set name
	c.Current_user_name = name

	// Get gatorconfig address
	gatorAddress, err := getConfigFilePath()
	if err != nil {
		return err
	}

	// Marshal data
	jsonData, err := json.Marshal(c)
	if err != nil {
		return err
	}

	// Write data
	err = os.WriteFile(gatorAddress, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
