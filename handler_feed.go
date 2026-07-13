package main

import (
	"context"
	"fmt"
	"time"

	"github.com/cocuum/aggreGator/internal/database"
	"github.com/google/uuid"
)

func handlerListFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetAllFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("Could not retrieve feeds: %w", err)
	}

	for i,feed := range feeds {
		printAllFeeds(i, feed)
	}

	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	if len(cmd.Args) != 2 {
		return fmt.Errorf("Usage: addfeed <name> <url> - name and url are required")
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:			uuid.New(),
		CreatedAt:	time.Now().UTC(),
		UpdatedAt:	time.Now().UTC(),
		Name: 		name,
		Url:		url,
		UserID:		user.ID,
	})
	if err != nil {
		return fmt.Errorf("Could not create feed: %w", err)
	}
	fmt.Println("Created feed successfully:")
	printFeed(feed)

	return nil
}

func  printFeed(feed database.Feed) {
	fmt.Printf("*ID:			%s\n",feed.ID)
	fmt.Printf("*Created at:	%v\n", feed.CreatedAt)
	fmt.Printf("*Updated at:	%v\n", feed.UpdatedAt)
	fmt.Printf("*Name:			%s\n",feed.Name)
	fmt.Printf("*URL:			%s\n",feed.Url)
	fmt.Printf("*User ID:		%s\n",feed.UserID)
}

func printAllFeeds (i int, feed database.GetAllFeedsRow) {
	fmt.Printf("======= %d ========\n", i+1)
	fmt.Printf("Feed Name:	%s\n", feed.FeedName)
	fmt.Printf("Feed URL:	%s\n", feed.Url)
	fmt.Printf("Username:	%s\n", feed.UserName)
	fmt.Println("==================")
}