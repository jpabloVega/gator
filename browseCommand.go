package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"strconv"
)

func browse(s *state, cmd command, user database.User) error {
	var limit int32
	if len(cmd.arguments) < 1 {
		limit = 2
	} else {
		temp, err := strconv.Atoi(cmd.arguments[0])
		limit = int32(temp)
		if err != nil {
			return err
		}
	}

	getPostParams := database.GetPostsForUserParams{
		ID:    user.ID,
		Limit: limit,
	}

	userPosts, err := s.db.GetPostsForUser(context.Background(), getPostParams)
	if err != nil {
		return err
	}

	for _, post := range userPosts {
		fmt.Printf("Title: %v\n", post.Title)
		fmt.Printf("Description: %v\n", post.Description.String)
		fmt.Printf("Publish Date: %v\n\n", post.PublishedAt)
	}
	return nil
}
