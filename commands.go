package main

import (
	"github.com/tonyserranodev/gator/internal/config"
	"github.com/tonyserranodev/gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	cmdNames map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	err := c.cmdNames[cmd.name](s, cmd)
	if err != nil {
		return err
	}

	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	if _, ok := c.cmdNames[name]; !ok {
		c.cmdNames[name] = f
	}
}
