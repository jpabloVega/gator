package main

import (
	"context"
	"errors"
	"fmt"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func addFeed(s *state, cmd command) error {
	// Check for expected arguments
	if len(cmd.arguments) < 2 {
		return errors.New("Correct syntax: go run . addfeed <feed name> <url>")
	}

	// Create context and get current user
	contx := context.Background()
	user, err := s.db.GetUser(contx, s.config.Current_user_name)
	if err != nil {
		return err
	}

	// Get values to pass to the db
	feedUUID := uuid.New()
	currTime := time.Now().UTC()
	name := cmd.arguments[0]
	url := cmd.arguments[1]
	user_id := user.ID

	// Create params to add to the db
	feedData := database.CreateFeedParams{
		ID:        feedUUID,
		CreatedAt: currTime,
		UpdatedAt: currTime,
		Name:      name,
		Url:       url,
		UserID:    user_id,
	}

	// Add to db
	feed, err := s.db.CreateFeed(contx, feedData)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", feed)
	return nil
}
