package main

import "fmt"

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func NewCommands() commands {
	registeredCommands := make(map[string]func(*state, command) error)
	return commands{registeredCommands}
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	val, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return fmt.Errorf("command '%s' is not registered", cmd.Name)
	}
	return val(s, cmd)
}
