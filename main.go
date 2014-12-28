package main

import (
	"fmt"
	"github.com/sinks/bench/command"
	"os"
	"time"
)

const dir_time_layout = "20060102"

var commands = []command.Command{
	{[]string{"new"}, NewHandler},
	{[]string{"status"}, StatusHandler},
	{[]string{"log"}, LogHandler},
}

func NewHandler() {
	fmt.Println("new day started")
}

func StatusHandler() {
	fmt.Println("today not started")
	fmt.Println("")
	fmt.Println("run: bench new")
}

func LogHandler() {
	fmt.Println("")
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
	fmt.Println("   log      create a new entry")
}

func new_day(base_path string) {
	time_now := time.Now()
	path := base_path + time_now.Format(dir_time_layout)
	err := os.MkdirAll(path, os.ModeDir)
	if err != nil {
		fmt.Println("Unable to create path")
	}
}
