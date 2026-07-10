package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/tonyserranodev/gator/internal/database"
	"log"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("must provide a username")
	}

	username := cmd.args[0]

	id := uuid.New()

	params := database.CreateUserParams{
		ID:        id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	}

	user, err := s.db.CreateUser(context.Background(), params)
	if err != nil {
		fmt.Printf("user already exists %s\n", user.Name)
		os.Exit(1)
	}

	s.cfg.CurrentUserName = user.Name
	err = s.cfg.SetUser()
	if err != nil {
		return err
	}

	fmt.Printf("user created: %s\n", username)

	log.Printf(`
	user details
	ID: %v
	CreatedAt: %v
	UpdatedAt: %v
	Name: %v

`, user.ID, user.CreatedAt, user.UpdatedAt, user.Name)

	return nil
}
