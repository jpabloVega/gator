package config

import (
	"errors"
)

func (c *commands) run(s *State, cmd command) error {
	// Check if command exists
	if funct, ok := c.command[cmd.name]; ok {
		// Run function
		err := funct(s, cmd)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("Invalid command")
	}
}

func (c *commands) register(name string, f func(*State, command) error) {
	c.command[name] = f
}
