package main

import (
	"gator/internal/config"
)

type command struct {
	name      string
	arguments []string
}

type state struct {
	config *config.Config
}

type commands struct {
	cmdFunc map[string]func(*state, command) error
}

func StartCommands() commands {
	return commands{
		cmdFunc: make(map[string]func(*state, command) error),
	}
}

func GetUserCommand(userCmds []string) command {
	var userArgs []string
	if len(userCmds) > 2 {
		userArgs = append(userArgs, userCmds[2:]...)
	}
	return command{
		name:      userCmds[1],
		arguments: userArgs,
	}
}
