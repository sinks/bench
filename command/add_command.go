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

const entry_time_format = "01-02 15:04"

type AddCommand struct {
}

func (ac *AddCommand) Handle() {
	var db *sql.DB
	before_handlers := []func() error{
		InitCheckHandler,
		func() error {
			var err error
			db, err = DbSetupHandler()
			return err
		},
	}

	TriggerHandlers(before_handlers)
	var entry *bench.Entry
	var err error
	if entry, err = AddHandler(db); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	after_handlers := []func() error{
		func() error {
			fmt.Println(entry.Time.Format(entry_time_format)+":", entry.Role, entry.Project, entry.Description)
			return nil
		},
	}
	TriggerHandlers(after_handlers)
}

func TriggerHandlers(handlers []func() error) {
	for _, f := range handlers {
		if err := f(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
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

func AddHandler(db *sql.DB) (*bench.Entry, error) {
	add_options := AddOptions{}
	flagSet := SetupFlags(&add_options)
	flagSet.Parse(os.Args[2:])
	add_options.Description = strings.Join(flagSet.Args(), " ")
	var tx *sql.Tx
	var err error
	if tx, err = db.Begin(); err != nil {
		return nil, err
	}
	var entry *bench.Entry
	entry, err = saveEntry(tx, add_options)
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return entry, nil
}

func saveEntry(tx *sql.Tx, add_options AddOptions) (*bench.Entry, error) {
	entry := &bench.Entry{
		time.Now(),
		add_options.Role,
		add_options.Project,
		add_options.Description,
	}
	err := entry.Save(tx)
	if err != nil {
		return nil, err
	}
	return entry, nil
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
