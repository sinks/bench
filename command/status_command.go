package command

import (
	"fmt"
	"github.com/sinks/bench/bench"
	"math"
	"os"
	"time"
)

type StatusCommand struct {
	db bench.BenchDatabase
}

const (
	display_time_format = "15:04:05"
	display_day_format  = "Monday 2006-01-02"
)

func (self *StatusCommand) Handle() {
	InitCheck()
	time_now := time.Now()
	time_start := BeginingOfDay(time_now)
	time_end := EndOfDay(time_now)
	results, err := bench.EntriesBetween(&self.db, time_start, time_end)
	if err != nil {
		fmt.Println("error getting entries")
		os.Exit(1)
	}

	fmt.Println("Today - " + time_now.Format(display_day_format))
	fmt.Println("")
	if results != nil {
		for i, entry := range results {
			fmt.Printf(
				"  [%d] %s %s@%s - %s\n",
				(i + 1),
				entry.Time.Format(display_time_format),
				entry.Role,
				entry.Project,
				entry.Description,
			)
		}
		fmt.Println("")
		duration := time_now.Sub(results[0].Time)
		fmt.Println("worked for hh:", math.Floor(duration.Hours()), "mm:", math.Floor(duration.Minutes()))
	} else {
		fmt.Println("nothing")
	}
}

func (self StatusCommand) Usage() {
}

func BeginingOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func EndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 99, t.Location())
}

func (sc *StatusCommand) Names() []string {
	return []string{"status"}
}
