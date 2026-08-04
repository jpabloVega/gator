package main

import (
	"context"
	"fmt"
)

func aggregator(s *state, cmd command) error {
	contx := context.Background()
	url := "https://www.wagslane.dev/index.xml"
	resp, err := s.client.FetchFeed(contx, url)
	if err != nil {
		return err
	}
	fmt.Println(resp)
	return nil
}
