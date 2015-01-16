package command

import (
	"flag"
	"fmt"
	"github.com/sinks/bench/bench"
	"os"
	"strings"
	"time"
)

const (
	entry_time_format = "01-02 15:04"
)

type AddCommand struct {
	db    bench.BenchDatabase
	flags *flag.FlagSet
}

func (self *AddCommand) Handle() {
	InitCheck()
	entry := self.flagParse()
	if err := self.db.Save(&entry); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Added")
}

func (self *AddCommand) flagParse() bench.Entry {
	var role string
	var project string
	var description string
	self.flags = flag.NewFlagSet("add", flag.ExitOnError)
	self.flags.Usage = func() { self.Usage() }
	self.flags.StringVar(&role, "role", "", "")
	self.flags.StringVar(&role, "r", "", "")
	self.flags.StringVar(&project, "project", "", "")
	self.flags.StringVar(&project, "p", "", "")
	self.flags.Parse(os.Args[2:])
	description = strings.Join(self.flags.Args(), " ")
	entry := bench.Entry{
		Role:        role,
		Project:     project,
		Description: description,
		Time:        time.Now(),
	}
	return entry
}

func (ac *AddCommand) Names() []string {
	return []string{"add"}
}

func (self AddCommand) Usage() {
	fmt.Println(
		"usage: bench add [-r <role>] [-p <project>] <description>\n",
		"\n",
		"options:\n",
		"   -r/--role      set the role\n",
		"   -p/--project   set the project",
	)
}
