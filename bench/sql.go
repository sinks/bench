package bench

import (
	"database/sql"
)

type SqlInserter interface {
	SqlInsert(tx *sql.Tx) error
}

type SqlUpdater interface {
	SqlUpdate(tx *sql.Tx) error
}

type SqlInserterUpdater interface {
	SqlInserter
	SqlUpdater
	SqlUpdateOrInsert(tx *sql.Tx) error
}
