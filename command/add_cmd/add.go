package add_cmd

import (
	"flag"
	"fmt"
	"github.com/sinks/bench/bench"
	"os"
	"strings"
	"time"
)

var (
	addFlags    *flag.FlagSet
	role        string
	project     string
	description string
)

func init() {
	addFlags = flag.NewFlagSet("add", flag.ExitOnError)
	addFlags.Usage = func() { usage() }
	addFlags.StringVar(&role, "role", "", "")
	addFlags.StringVar(&role, "r", "", "")
	addFlags.StringVar(&project, "project", "", "")
	addFlags.StringVar(&project, "p", "", "")

}

func flagParse() bench.Entry {
	addFlags.Parse(os.Args[2:])
	description = strings.Join(addFlags.Args(), " ")
	entry := bench.Entry{
		Role:        role,
		Project:     project,
		Description: description,
		Time:        time.Now(),
	}
	return entry
}

func usage() {
	fmt.Println(
		"usage: bench add [-r <role>] [-p <project>] <description>\n",
		"\n",
		"options:\n",
		"   -r/--role      set the role\n",
		"   -p/--project   set the project",
	)
}
