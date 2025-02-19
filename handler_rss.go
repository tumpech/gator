package main

import (
	"context"
	"fmt"
	"time"

	"github.com/tumpech/gator/internal/database"
)

func printFeed(feed database.Feed, user database.User) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* User:          %s\n", user.Name)
}

func handlerAgg(s *state, cmd command) error {
	feedURL := "https://www.wagslane.dev/index.xml"
	rssFeed, err := fetchFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("error fetching URL: %w", err)
	}
	fmt.Printf("RSSFeed: %+v\n", rssFeed)
	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <Name> <URL>", cmd.Name)
	}
	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error getting current user: %w", err)
	}

	arg := database.CreateFeedParams{
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    currentUser.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), arg)
	if err != nil {
		return fmt.Errorf("error creating feed in DB: %w", err)
	}

	fmt.Println("Feed created successfully:")
	printFeed(feed, currentUser)
	fmt.Println()
	fmt.Println("=====================================")

	return nil
}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.ListFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error gathering feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds in the DB.")
		return nil
	}

	fmt.Printf("Found %d feeds:\n", len(feeds))

	for _, feed := range feeds {
		user, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("error gahtering username by ID: %w", err)
		}
		printFeed(feed, user)
		fmt.Println("=====================================")
	}
	return nil
}
