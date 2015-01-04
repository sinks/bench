package bench

import (
	"database/sql"
)

type Document struct {
	Entries []Entry `json:"entries"`
}

type DocumentSQL struct {
	id   sql.NullInt64
	Data Document
}

func (ds *DocumentSQL) SaveUpdate(tx *sql.Tx) error {
	if ds.id.Valid == true {
		// update document
	} else {
		// create document
	}
	return nil
}

func FindDocument(doc *Document, id int64) error {
	return nil
}
