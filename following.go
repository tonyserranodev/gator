package main

import (
	"context"
	"fmt"

	"github.com/tonyserranodev/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	feed_follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error getting feed follows: %s", err)
	}

	for _, f := range feed_follows {
		fmt.Printf("%s's feeds:\n", user.Name)
		fmt.Printf("  %s\n", f.FeedName)
	}

	return nil
}
