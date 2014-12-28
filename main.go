package main

import (
	"fmt"
	"github.com/sinks/bench/command"
	"os"
)

var commands = []command.Command{
	{[]string{"new"}, command.NewHandler},
	{[]string{"add"}, command.AddHandler},
	{[]string{"status"}, StatusHandler},
	{[]string{"commit"}, CommitHandler},
	{[]string{"history"}, HistoryHandler},
	{[]string{"checkout"}, CheckoutHandler},
}

func CheckoutHandler() {
	fmt.Println("")
}

func HistoryHandler() {
	fmt.Println("")
}

func StatusHandler() {
	fmt.Println("bench not created")
	fmt.Println("")
	fmt.Println("run: bench new")
}

func CommitHandler() {
	fmt.Println("log handler")
}

func Commands(command_name string) (*command.Command, error) {
	for _, command := range commands {
		if command.HandlesName(command_name) {
			return &command, nil
		}
	}
	return nil, command.CommandNotValidError(command_name)
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