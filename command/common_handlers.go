package command

import (
	"errors"
	"github.com/sinks/bench/bench"
)

var error_init_not_run = errors.New(`bench not setup

run: bench init`)

func InitCheckHandler() error {
	if !bench.BenchDirExists() {
		return error_init_not_run
	}
	return nil
}
