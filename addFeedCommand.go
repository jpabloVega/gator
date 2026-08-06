package main

import (
	"context"
	"errors"
	"fmt"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func addFeed(s *state, cmd command, user database.User) error {
	// Check for expected arguments
	if len(cmd.arguments) < 2 {
		return errors.New("Correct syntax: go run . addfeed <feed name> <url>")
	}

	// Create context
	contx := context.Background()

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

	// Add feed to db
	feed, err := s.db.CreateFeed(contx, feedData)
	if err != nil {
		return err
	}

	// Add feed follow
	feedFollowData := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user_id,
		FeedID:    feed.ID,
	}
	_, err = s.db.CreateFeedFollow(contx, feedFollowData)
	if err != nil {
		return err
	}

	fmt.Printf("%s added feed %v\n", user.Name, feed.Name)
	fmt.Printf("%+v\n", feed)
	return nil
}
