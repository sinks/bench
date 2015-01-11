package bench

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type BenchDatabase struct {
	db *sql.DB
}

// Save will save any number of documents who support
// SqlInserter in a single transaction
func (self *BenchDatabase) Save(docs ...SqlInserter) error {
	if self.db == nil {
		err := self.Connect()
		if err != nil {
			return err
		}
	}
	tx, err := self.db.Begin()
	if err != nil {
		return err
	}
	for _, doc := range docs {
		if err = doc.SqlInsert(tx); err != nil {
			if rb_err := tx.Rollback(); rb_err != nil {
				return rb_err
			}
			return err
		}
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (self *BenchDatabase) Create() error {
	if self.db == nil {
		err := self.Connect()
		if err != nil {
			return err
		}
	}
	tx, err := self.db.Begin()
	if err != nil {
		return err
	}
	if err = schemaRun(tx); err != nil {
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

func (self *BenchDatabase) Connect() error {
	var err error
	self.db, err = sql.Open("sqlite3", BenchDBPath)
	if err != nil {
		return err
	}
	if err = self.db.Ping(); err != nil {
		return err
	}
	return nil
}
