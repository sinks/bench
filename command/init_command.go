package command

import (
	"fmt"
	"github.com/sinks/bench/bench"
	"os"
)

const (
	dir_perms = 0755
)

type InitCommand struct {
	db bench.BenchDatabase
}

func (self *InitCommand) Handle() {
	if err := createBenchDir(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := self.db.Create(); err != nil {
	}
	fmt.Println("bench created")
}

func (self *InitCommand) Names() []string {
	return []string{"init"}
}

func createBenchDir() error {
	err := os.MkdirAll(bench.BenchDir, dir_perms)
	if err != nil {
		return fmt.Errorf("failed to create %s", bench.BenchDir)
	}
	return nil
}
