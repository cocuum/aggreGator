package main

import (
	"context"
	"fmt"
	"log"
	"time"
	
	"github.com/cocuum/aggreGator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Usage: agg - one argument required")
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("Could not set reqs intervals: %w\n", err)
	}

	fmt.Printf("Collecting feeds every %s ...\n", timeBetweenReqs)

	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) {
	nextFeed,err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Printf("Could not get next feed to fetch: %v\n", err)
		return
	}

	log.Println("Found feed to fetch ...")
	scrapeFeed(s.db, nextFeed)
}

func scrapeFeed(db *database.Queries, feed database.Feed) {
	_, err := db.MarkFeedFetched(context.Background(), feed.ID)

	if err != nil {
		log.Printf("Could not mark feed as fetched: %v\n", err)
		return
	}

	feedData, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("Could not fetch feed %s: %v\n", feed.Name,err)
		return
	}

	for _,feed := range feedData.Channel.Item {
		fmt.Printf("Found post: %s\n", feed.Title)
	}

	log.Printf("Feed %s collected, %v posts found", feed.Name, len(feedData.Channel.Item))
}