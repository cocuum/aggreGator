package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/cocuum/aggreGator/internal/database"
	"github.com/google/uuid"
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

	for _, post := range feedData.Channel.Item {
		
		// Convert description format for database
		var description sql.NullString
		if post.Description != "" {
			description = sql.NullString{String: post.Description, Valid: true}
		} else {
			description = sql.NullString{String: "", Valid: false}
		}

		//Convert published date format for database
		var publishedAt sql.NullTime
		if post.PubDate != "" {
			
			t,_ := time.Parse(time.DateTime, post.PubDate)
			publishedAt = sql.NullTime{Time: t, Valid: true}
		} else {
			t,_ := time.Parse(time.DateTime, "0000-00-00 00:00:00")
			publishedAt = sql.NullTime{Time: t, Valid: false}
		}

		_, err := db.CreatePost(context.Background(), database.CreatePostParams{
			ID:				uuid.New(),
			CreatedAt:		time.Now().UTC(),
			UpdatedAt:		time.Now().UTC(),
			Title:			post.Title,
			Url:			post.Link,
			Description:	description,
			PublishedAt:	publishedAt,
			FeedID:			feed.ID,
		})

		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("Could not create post %v\n", err)
			return
		}
	}

	log.Printf("Feed %s collected, %v posts found", feed.Name, len(feedData.Channel.Item))
}