package command

import (
	"flag"
	"fmt"
	"github.com/sinks/bench/bench"
	"os"
	"strings"
	"time"
)

const entry_time_format = "01-02 15:04"

type AddCommand struct {
	db bench.BenchDatabase
}

func (ac *AddCommand) Handle() {
	InitCheckHandler()
	add_options := AddOptions{}
	flagSet := SetupFlags(&add_options)
	flagSet.Parse(os.Args[2:])
	add_options.Description = strings.Join(flagSet.Args(), " ")
	err := ac.db.Save(add_options.ToEntry())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (ac *AddCommand) Names() []string {
	return []string{"add"}
}

type AddOptions struct {
	Role        string
	Project     string
	Description string
}

func (ao AddOptions) ToEntry() *bench.Entry {
	return &bench.Entry{
		time.Now(),
		ao.Role,
		ao.Project,
		ao.Description,
	}
}

func SetupFlags(add_options *AddOptions) *flag.FlagSet {
	flagSet := flag.NewFlagSet("new", flag.ExitOnError)
	flagSet.Usage = func() { PrintUsage() }
	flagSet.StringVar(&add_options.Role, "role", "", "")
	flagSet.StringVar(&add_options.Role, "r", "", "")
	flagSet.StringVar(&add_options.Project, "project", "", "")
	flagSet.StringVar(&add_options.Project, "p", "", "")
	return flagSet
}

func Usage() string {
	return fmt.Sprintln(
		"usage: bench add [-r <role>] [-p <project>] <description>\n",
		"\n",
		"options:\n",
		"   -r/--role      set the role\n",
		"   -p/--project   set the project",
	)
}

func PrintUsage() {
	fmt.Println(Usage())
}
