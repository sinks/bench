package command

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/sinks/bench/bench"
	"os"
	"strings"
	"time"
)

type NewCommand Command

func NewAddCommand() *Command {
	var db *sql.DB
	before_handlers := []func() error{
		InitCheckHandler,
		DbSetupHandler(db),
	}
	handler := AddHandler(db)
	after_handlers := []func() error{}

	command := &Command{
		[]string{"add"},
		handler,
		before_handlers,
		after_handlers,
	}
	return command
}

type AddOptions struct {
	Role        string
	Project     string
	Description string
}

func AddHandler(db *sql.DB) func() error {
	return func() error {
		DbSetupHandler(db)
		fmt.Println(db)
		add_options := AddOptions{}
		flagSet := SetupFlags(&add_options)
		flagSet.Parse(os.Args[2:])
		add_options.Description = strings.Join(flagSet.Args(), " ")
		var tx *sql.Tx
		var err error
		fmt.Println(db)
		if tx, err = db.Begin(); err != nil {
			return err
		}
		saveEntry(tx, add_options)
		if err = tx.Commit(); err != nil {
			return err
		}
		return nil
	}
}

func saveEntry(tx *sql.Tx, add_options AddOptions) error {
	entry := &bench.Entry{
		time.Now(),
		add_options.Role,
		add_options.Project,
		add_options.Description,
	}

	err := entry.Save(tx)
	if err != nil {
		return err
	}

	return nil
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
