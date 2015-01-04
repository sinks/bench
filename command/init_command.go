package command

import (
	"database/sql"
	"fmt"
	"github.com/sinks/bench/bench"
	"os"
)

type InitCommand struct {
}

func (ic *InitCommand) Handle() {
	var db *sql.DB
	InitHandler(db)
	after_handlers := []func() error{success}
	TriggerHandlers(after_handlers)
}

func (ic *InitCommand) Names() []string {
	return []string{"init"}
}

const dir_perms = 0755

type InitCommandCreatePathError string

func (s InitCommandCreatePathError) Error() string {
	return fmt.Sprintf("failed to create %s", string(s))
}

func InitHandler(db *sql.DB) error {
	err := initBenchDir(bench.BenchPath())
	if err != nil {
		return err
	}
	err = initBenchDatabase(db)
	if err != nil {
		return err
	}
	return nil
}

func initBenchDir(base_path string) error {
	err := os.MkdirAll(base_path, dir_perms)
	if err != nil {
		return InitCommandCreatePathError(base_path)
	}
	return nil
}

func initBenchDatabase(db *sql.DB) error {
	var err error
	db, err = sql.Open("sqlite3", bench.DbPath())
	if err != nil {
		return err
	}
	return bench.CreateDatabase(db)
}

func success() error {
	fmt.Println("bench created")
	return nil
}
