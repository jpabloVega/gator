package main

import (
	"context"
	"errors"
	"fmt"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func registerUser(s *state, cmd command) error {
	// Check if name was passed
	if len(cmd.arguments) == 0 {
		return errors.New("No name provided")
	}

	// Create new user in the db
	// Get context data
	contx := context.Background()
	userUUID := uuid.New()
	currTime := time.Now().UTC()
	userName := cmd.arguments[0]
	userParams := database.CreateUserParams{
		ID:        userUUID,
		CreatedAt: currTime,
		UpdatedAt: currTime,
		Name:      userName,
	}

	// Check if user exists
	_, err := s.db.GetUser(contx, userName)
	if err == nil {
		return errors.New("User already exists")
	}

	// Create user
	userData, err := s.db.CreateUser(contx, userParams)
	if err != nil {
		return err
	}

	// Set config user to username
	err = s.config.SetUser(userName)
	if err != nil {
		return err
	}

	// Print user information
	fmt.Printf("%s was registered successfuly\n", userName)
	fmt.Printf("%+v\n", userData)
	return nil
}
