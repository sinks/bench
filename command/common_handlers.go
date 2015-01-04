package command

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sinks/bench/bench"
)

var error_db_not_setup = errors.New(`database not setup

run: bench init`)

var error_init_not_run = errors.New(`bench not setup

run: bench init`)

func InitCheckHandler() error {
	if !bench.BenchDirExists() {
		return error_init_not_run
	}
	return nil
}

func DbSetupHandler(db *sql.DB) func() error {
	return func() error {
		var err error
		if !bench.DbExists() {
			return error_db_not_setup
		}
		db, err = sql.Open("sqlite3", bench.DbPath())
		if err != nil {
			return err
		}
		err = db.Ping()
		return err
	}
}
