package main

import (
	"context"
	"fmt"
)

func resetTable(s *state, cmd command) error {
	contx := context.Background()
	err := s.db.DeleteUsers(contx)
	if err != nil {
		return err
	}
	fmt.Println("Users deleted")
	return nil
}
