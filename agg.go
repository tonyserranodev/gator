package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/tonyserranodev/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("must provide a timeBetweenReqs argument")
	}

	timeBetweenReqs := cmd.args[0]

	t, err := time.ParseDuration(timeBetweenReqs)
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %s\n", timeBetweenReqs)

	ticker := time.NewTicker(t)

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	params := database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt:     time.Now(),
		ID:            nextFeed.ID,
	}

	err = s.db.MarkFeedFetched(context.Background(), params)
	if err != nil {
		return err
	}

	feed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return err
	}

	for _, item := range feed.Channel.Item {

		pubDate, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			return err
		}

		params := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: pubDate,
			FeedID:      nextFeed.ID,
		}

		s.db.CreatePost(context.Background(), params)
		fmt.Println("Post Saved!")
		fmt.Printf("Title: %v\n", item.Title)
	}

	return nil
}
