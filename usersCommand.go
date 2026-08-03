package main

import (
	"context"
	"fmt"
)

func getUsers(s *state, cmd command) error {
	// Get users
	contx := context.Background()
	usersData, err := s.db.GetUsers(contx)
	if err != nil {
		return err
	}

	// List users
	currName := s.config.Current_user_name
	fmt.Println(currName)
	for _, userData := range usersData {
		fmt.Printf("* %s", userData.Name)
		if userData.Name == currName {
			fmt.Print(" (current)")
		}
		fmt.Println()
	}
	return nil
}
