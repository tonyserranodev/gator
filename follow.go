package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/tonyserranodev/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("must provide the url of a feed to follow")
	}

	url := cmd.args[0]

	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error getting feed: %s", err)
	}

	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	feed_follow, err := s.db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error creating feed follow: %s", err)
	}

	fmt.Printf("Feed name: %s\n", feed_follow.FeedName)
	fmt.Printf("User name: %s\n", feed_follow.UserName)

	return nil
}
