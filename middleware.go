package main

import (
	"context"
	"gator/internal/database"
)

func mwLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.config.Current_user_name)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}
