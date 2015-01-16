package command

import (
	"fmt"
)

type Command interface {
	Handle()
	Usage()
	Names() []string
}

type CommandNotValidError string

func (c CommandNotValidError) Error() string {
	return fmt.Sprintf("Command %s not found", string(c))
}
