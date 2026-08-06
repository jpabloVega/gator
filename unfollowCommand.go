package main

import (
	"context"
	"errors"
	"fmt"
	"gator/internal/database"
)

func unfollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) < 1 {
		return errors.New("Correct syntax: go run . unfollow <feed url>")
	}

	feedData, err := s.db.GetFeedFromUrl(context.Background(), cmd.arguments[0])
	if err != nil {
		return err
	}

	deleteParams := database.DeleteFollowParams{
		UserID: user.ID,
		FeedID: feedData.ID,
	}

	err = s.db.DeleteFollow(context.Background(), deleteParams)
	if err != nil {
		return err
	}

	fmt.Printf("%s unfollowed", cmd.arguments[0])
	return nil
}
