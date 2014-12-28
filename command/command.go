package command

import (
	"fmt"
)

type Command struct {
	Names  []string
	Handle func()
}

func (c *Command) HandlesName(command string) bool {
	for _, name := range c.Names {
		if name == command {
			return true
		}
	}
	return false
}

type CommandNotValidError string

func (c CommandNotValidError) Error() string {
	return fmt.Sprintf("Command %s not found", string(c))
}
