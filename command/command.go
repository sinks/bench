package command

import (
	"github.com/sinks/bench/bench"
	"github.com/sinks/bench/command/add_cmd"
	"github.com/sinks/bench/command/init_cmd"
	"github.com/sinks/bench/command/status_cmd"
)

var (
	db = bench.BenchDatabase{}
)

type Command interface {
	Handle()
	Usage()
	Names() []string
}

func NewInitCommand() Command {
	return &init_cmd.Command{DB: db}
}

func NewAddCommand() Command {
	return &add_cmd.Command{DB: db}
}

func NewStatusCommand() Command {
	return &status_cmd.Command{DB: db}
}
