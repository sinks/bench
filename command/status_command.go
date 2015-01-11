package command

import (
	"github.com/sinks/bench/bench"
)

type StatusCommand struct {
	db *bench.BenchDatabase
}

func (sc *StatusCommand) Handle() {
	InitCheckHandler()
}

func (sc *StatusCommand) Names() []string {
	return []string{"status"}
}
