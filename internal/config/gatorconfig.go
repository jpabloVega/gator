package config

import (
	"os"
)

func getConfigFilePath() (string, error) {
	gatorConfig, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	gatorConfig += "/.gatorconfig.json"
	return gatorConfig, nil
}
