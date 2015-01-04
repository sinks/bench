package bench

import (
	"database/sql"
)

var document_schema = `
CREATE TABLE IF NOT EXISTS document (
	id INTEGER PRIMARY KEY
);
`

var entries_schema = `
CREATE TABLE IF NOT EXISTS entries (
	id INTEGER PRIMARY KEY,
	time INTEGER,
	role STRING,
	project STRING,
	description STRING
);
`

func CreateDatabase(db *sql.DB) error {
	var err error
	if err = db.Ping(); err != nil {
		return err
	}
	var tx *sql.Tx
	if tx, err = db.Begin(); err != nil {
		return err
	}
	if err = schemaCreate(tx); err != nil {
		if rb_err := tx.Rollback(); rb_err != nil {
			return rb_err
		}
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func schemaCreate(tx *sql.Tx) error {
	_, err := tx.Exec(document_schema)
	if err != nil {
		return err
	}
	_, err = tx.Exec(entries_schema)
	if err != nil {
		return err
	}
	return err
}
