package main

import (
	"errors"
)

func (c *commands) run(s *state, cmd command) error {
	// Check if command exists
	if funct, ok := c.cmdFunc[cmd.name]; ok {
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

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmdFunc[name] = f
}
