package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tonyserranodev/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading gatorconfig: %v", err)
	}

	st := &state{
		cfg: &cfg,
	}

	cmds := commands{
		cmdNames: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		fmt.Println("must provide the name of a command")
		os.Exit(1)
	}

	userArgs := os.Args

	cmdName := userArgs[1]

	cmdArgs := userArgs[2:]

	err = cmds.run(st, command{name: cmdName, args: cmdArgs})
	if err != nil {
		fmt.Println(err)
	}
}
