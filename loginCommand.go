package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	// Check the existance of arguments
	if len(cmd.arguments) == 0 {
		return errors.New("No arguments passed")
	}

	username := cmd.arguments[0]
	// Check if user exists in the db
	contx := context.Background()
	_, err := s.db.GetUser(contx, username)
	if err != nil {
		return errors.New("User doesnt exists")
	}

	// Set the username
	err = s.config.SetUser(username)
	if err != nil {
		return err
	}

	// Report the new user
	fmt.Printf("Added user: %s\n", username)

	return nil
}
