package bench

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"strings"
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

func (self *BenchDatabase) Query(query string, params ...interface{}) error {
	if self.db == nil {
		err := self.Connect()
		if err != nil {
			return err
		}
	}

	rows, _ := self.db.Query(query, params...)
	for rows.Next() {

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

func (self *BenchDatabase) Select(sql string, args ...interface{}) (*sql.Rows, error) {
	if self.db == nil {
		err := self.Connect()
		if err != nil {
			return nil, err
		}
	}
	result, err := self.db.Query(sql, args...)
	return result, err
}

func Select(fields []string) string {
	return fmt.Sprintf("SELECT %s", strings.Join(fields, ", "))
}

func From(table_name string) string {
	return fmt.Sprintf("FROM %s", table_name)
}

func Where(where ...string) string {
	return fmt.Sprintf("WHERE %s", where)
}

func Greater(field string, value int) string {
	return fmt.Sprintf("%s > %d", field, value)
}

func GreaterEqual(field string, value int) string {
	return fmt.Sprintf("%s >= %d", field, value)
}

func Lesser(field string, value int) string {
	return fmt.Sprintf("%s < %d", field, value)
}

func LesserEqual(field string, value int) string {
	return fmt.Sprintf("%s <= %d", field, value)
}
