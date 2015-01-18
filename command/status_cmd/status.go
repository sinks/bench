package status_cmd

import (
	"fmt"
	"github.com/sinks/bench/bench"
	"time"
)

func usage() {
}

func header(time time.Time) {
	fmt.Println("Today - " + time.Format(display_day_format))
	fmt.Println("")
}

func showEntries(entries []bench.Entry) {
	for i, entry := range entries {
		fmt.Println("  ", (i + 1), entry)
	}
}

func BeginingOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func EndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 99, t.Location())
}
