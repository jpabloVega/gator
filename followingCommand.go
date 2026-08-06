package main

import (
	"context"
	"gator/internal/database"

	"fmt"
)

func following(s *state, cmd command, user database.User) error {
	// Get feeds from user
	contx := context.Background()
	userFeeds, err := s.db.GetFeedFollowsForUser(contx, user.Name)
	if err != nil {
		return err
	}

	// Print results
	fmt.Printf("%s is following:\n", user.Name)
	for _, feed := range userFeeds {
		fmt.Printf(" * %s\n", feed.FeedName)
	}
	return nil
}
