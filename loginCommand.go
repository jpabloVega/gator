package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	// Check the existance of arguments
	if len(cmd.arguments) == 0 {
		return errors.New("No arguments passed")
	}

	// Set the username
	username := cmd.arguments[0]
	err := s.config.SetUser(username)
	if err != nil {
		return err
	}

	// Report the new user
	fmt.Printf("Added user: %s\n", username)

	return nil
}
