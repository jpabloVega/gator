package main

import (
	"context"
	"fmt"
)

func feeds(s *state, cmd command) error {
	contx := context.Background()
	feeds, err := s.db.GetFeeds(contx)
	if err != nil {
		return err
	}
	for _, feed := range feeds {
		fmt.Printf(" * %s\n", feed.Name)
		fmt.Printf(" * %s\n", feed.Url)
		fmt.Printf(" * %s\n", feed.UserName)
		fmt.Println()
	}
	return nil
}
