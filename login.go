package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("must provide a username")
	}

	username := cmd.args[0]

	s.cfg.CurrentUserName = username

	err := s.cfg.SetUser()
	if err != nil {
		return err
	}

	fmt.Printf("user: %s has been set\n", username)

	return nil
}
