package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	feedURL := "https://www.wagslane.dev/index.xml"
	rssFeed, err := fetchFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("error fetching URL: %w", err)
	}
	fmt.Printf("RSSFeed: %+v\n", rssFeed)
	return nil
}
