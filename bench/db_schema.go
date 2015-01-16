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
	time datetime,
	role STRING,
	project STRING,
	description STRING
);
`

func schemaRun(tx *sql.Tx) error {
	_, err := tx.Exec(document_schema)
	if err != nil {
		return err
	}
	_, err = tx.Exec(entries_schema)
	if err != nil {
		return err
	}
	return nil
}
