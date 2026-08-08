package main

import (
	"context"
	"errors"
	"fmt"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func follow(s *state, cmd command, user database.User) error {
	// Check for arguments
	if len(cmd.arguments) < 1 {
		return errors.New("Correct syntaxis: gator follow <url>")
	}
	contx := context.Background()

	// Get feed data
	feedData, err := s.db.GetFeedFromUrl(contx, cmd.arguments[0])
	if err != nil {
		return err
	}

	// Define feed follow params
	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feedData.ID,
	}

	// Add to database
	feedFollowData, err := s.db.CreateFeedFollow(contx, feedFollowParams)
	if err != nil {
		return err
	}

	fmt.Printf("%s is now following %s\n", feedFollowData.UserName, feedFollowData.FeedName)
	return nil
}
