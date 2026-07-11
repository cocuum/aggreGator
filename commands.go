package main

import "errors"

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCMDS map[string]func(*state, command) error
}

// registers a new handler function for a command name
func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCMDS[name] = f
}

// runs a given command with the provided state if it exists
func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.registeredCMDS[cmd.Name]
	if !ok {
		return errors.New("Unknown command:" + cmd.Name)
	}
	return handler(s, cmd)
}
