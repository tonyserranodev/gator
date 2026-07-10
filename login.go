package main

import (
	"context"
	"errors"
	"fmt"
	"os"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("must provide a username")
	}

	username := cmd.args[0]

	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		fmt.Printf("username %s not found\n", username)
		os.Exit(1)
	}

	s.cfg.CurrentUserName = user.Name

	err = s.cfg.SetUser()
	if err != nil {
		return err
	}

	fmt.Printf("user: %s has been set\n", username)

	return nil
}
