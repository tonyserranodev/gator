package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		return fmt.Errorf("feeds command takes no arguments")
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error getting feeds: %v", err)
	}

	for _, f := range feeds {
		user, err := s.db.GetUserByID(context.Background(), f.UserID)
		if err != nil {
			return fmt.Errorf("error getting user: %v", err)
		}

		fmt.Printf("Feed name: %s\n", f.Name)
		fmt.Printf("Url: %s\n", f.Url)
		fmt.Printf("Url: %s\n", f.Url)
		fmt.Printf("Feed created by: %s\n", user.Name)

	}

	return nil
}
