package bench

import (
	"database/sql"
	"fmt"
	"time"
)

const (
	insert_sql       = "INSERT INTO entries (time, role, project, description) VALUES (?, ?, ?, ?);"
	time_between_sql = "SELECT id, time, role, project, description FROM ENTRIES WHERE time >= ? AND time <= ?;"
)

const (
	entry_time_format = "2006-01-02 15:04"
	entry_format      = "%s: %s@%s - %s"
)

type Entry struct {
	Time        time.Time `json:"time"`
	Role        string    `json:"role"`
	Project     string    `json:"project"`
	Description string    `json:"description"`
}

func (e *Entry) SqlInsert(tx *sql.Tx) error {
	time_utc := e.Time.UTC()
	_, err := tx.Exec(insert_sql, time_utc, e.Role, e.Project, e.Description)
	if err != nil {
		return err
	}
	return nil
}

func EntriesBetween(db *BenchDatabase, start time.Time, end time.Time) ([]Entry, error) {
	rows, err := db.Select(time_between_sql, start, end)
	if err != nil {
		return nil, err
	}

	var results []Entry
	for rows.Next() {
		var result Entry
		var id int64
		var utc_time time.Time
		err = rows.Scan(&id, &utc_time, &result.Role, &result.Project, &result.Description)
		result.Time = utc_time.Local()
		if err == nil {
			results = append(results, result)
		}
	}
	return results, nil
}

func (entry Entry) String() string {
	return fmt.Sprintf(
		entry_format,
		entry.Time.Format(entry_time_format),
		entry.Role,
		entry.Project,
		entry.Description,
	)
}
