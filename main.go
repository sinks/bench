package main

import (
	"errors"
	"fmt"
	"github.com/sinks/bench/command"
	"os"
)

var commands = []command.Command{
	command.NewInitCommand(),
	command.NewAddCommand(),
	command.NewStatusCommand(),
}

func Commands(command_name string) (command.Command, error) {
	for _, command := range commands {
		if handlesName(command_name, command.Names()) {
			return command, nil
		}
	}
	return nil, errors.New("command not valid")
}

func handlesName(command string, names []string) bool {
	for _, name := range names {
		if name == command {
			return true
		}
	}
	return false
}

func main() {
	if len(os.Args) > 1 {
		command, err := Commands(os.Args[1])
		if err == nil {
			command.Handle()
		} else {
			Usage()
		}
	} else {
		Usage()
	}
}

func Usage() {
	fmt.Println("usage: bench [--help] <command> [<args>]")
	fmt.Println("")
	fmt.Println("The most commonly used commands are:")
	fmt.Println("   new      start a new day")
	fmt.Println("   status   show current times")
	fmt.Println("   commit   create a new entry")
}
