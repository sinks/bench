package command

import (
	"fmt"
	"os"
)

type Command interface {
	Handle()
}

type Command struct {
	Names        []string
	handle       func() error
	beforeHandle []func() error
	afterHandle  []func() error
}

func (c *Command) Handle() {
	if err := c.triggerBeforeHandle(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := c.handle(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := c.triggerAfterHandle(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (c *Command) triggerBeforeHandle() error {
	for _, f := range c.beforeHandle {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}

func (c *Command) triggerAfterHandle() error {
	for _, f := range c.afterHandle {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
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
