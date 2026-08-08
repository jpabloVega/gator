package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"gator/internal/api"
	"gator/internal/database"
	"strings"
	"time"

	"github.com/google/uuid"
)

func aggregator(s *state, cmd command) error {
	if len(cmd.arguments) < 1 {
		return errors.New("Correct syntax: go run . agg <time duration examp: 1m, 2h>")
	}

	time_between_reqs, err := time.ParseDuration(cmd.arguments[0])
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %v\n", time_between_reqs)
	ticker := time.NewTicker(time_between_reqs)

	for ; ; <-ticker.C {
		_, err := scrapeFeeds(s)
		if err != nil {
			return err
		}
	}
}

func scrapeFeeds(s *state) (*api.XMLResponse, error) {
	// Get the next feed to fetch
	feedToFetch, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return &api.XMLResponse{}, err
	}

	// Mark feed as fetched
	markedFeed, err := s.db.MarkFeedFetched(context.Background(), feedToFetch.ID)
	if err != nil {
		return &api.XMLResponse{}, err
	}

	// Get feed from url
	resp, err := s.client.FetchFeed(context.Background(), markedFeed.Url)
	if err != nil {
		return &api.XMLResponse{}, err
	}

	for _, feed := range resp.Channel.Item {
		publishedDate, err := time.Parse(time.RFC1123Z, feed.PubDate)
		if err != nil {
			return &api.XMLResponse{}, err
		}
		createPostParams := database.CreatePostParams{
			ID:          uuid.New(),
			Title:       feed.Title,
			Url:         feed.Link,
			Description: sql.NullString{String: feed.Description, Valid: true},
			PublishedAt: publishedDate,
			FeedID:      uuid.NullUUID{UUID: markedFeed.ID, Valid: true},
		}
		_, err = s.db.CreatePost(context.Background(), createPostParams)
		if err != nil {
			if strings.Contains(err.Error(), "unique constraint") {
				continue
			}
			return &api.XMLResponse{}, err
		}
	}

	return resp, nil
}
