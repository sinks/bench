package init_cmd

import (
	"fmt"
	"github.com/sinks/bench/bench"
	"os"
)

const (
	dir_perms = 0755
)

type Command struct {
	DB bench.BenchDatabase
}

func (self *Command) Handle() {
	if err := createBenchDir(); err != nil {
		fmt.Println("unable to create bench dir")
		fmt.Println(err)
		os.Exit(1)
	}
	if err := self.DB.Create(); err != nil {
		fmt.Println("unable to create database")
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("bench created")
}

func (self *Command) Names() []string {
	return []string{"init"}
}

func (self Command) Usage() {
}

func createBenchDir() error {
	if err := os.MkdirAll(bench.BenchDir, dir_perms); err != nil {
		return err
	}
	return nil
}
