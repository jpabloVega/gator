package config

import (
	"errors"
	"fmt"
)

func HandlerLogin(s *State, cmd command) error {
	// Check the existance of arguments
	if len(cmd.arguments) == 0 {
		return errors.New("No arguments passed")
	}

	// Set the username
	username := cmd.arguments[0]
	err := s.Config.SetUser(username)
	if err != nil {
		return err
	}

	// Report the new user
	fmt.Printf("Added user: %s\n", username)

	return nil
}
