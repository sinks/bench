package bench

import (
	"database/sql"
	"time"
)

const insert_sql = "INSERT INTO entries (time, role, project, description) VALUES (?, ?, ?, ?);"

type Entry struct {
	Time        time.Time `json:"time"`
	Role        string    `json:"role"`
	Project     string    `json:"project"`
	Description string    `json:"description"`
}

func (e *Entry) Save(tx *sql.Tx) error {
	_, err := tx.Exec(insert_sql, e.Time, e.Role, e.Project, e.Description)
	if err != nil {
		return err
	}
	return nil
}

type ByEntryTimes []Entry

func (b ByEntryTimes) Len() int           { return len(b) }
func (b ByEntryTimes) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByEntryTimes) Less(i, j int) bool { return b[i].Time.After(b[j].Time) }
