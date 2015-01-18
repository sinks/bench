package add_cmd

import (
	"fmt"
	"github.com/sinks/bench/bench"
	"github.com/sinks/bench/command/checks"
	"os"
)

type Command struct {
	DB bench.BenchDatabase
}

func (self *Command) Handle() {
	checks.InitCheck()
	entry := flagParse()
	if err := self.DB.Save(&entry); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Added", entry)
}

func (ac Command) Names() []string {
	return []string{"add"}
}

func (ac Command) Usage() {
	usage()
}
