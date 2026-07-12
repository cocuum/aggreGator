package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) > 0 {
		return fmt.Errorf("Usage: agg - no arguments are required")
	}

	url := "https://www.wagslane.dev/index.xml"

	rssFeed, err := fetchFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}

	fmt.Printf("rssFeed: %+v\n", rssFeed)
	return nil
}
