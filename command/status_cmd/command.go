package status_cmd

import (
	"fmt"
	"github.com/sinks/bench/bench"
	"github.com/sinks/bench/command/checks"
	"math"
	"os"
	"time"
)

type Command struct {
	DB bench.BenchDatabase
}

const (
	display_time_format = "15:04:05"
	display_day_format  = "Monday 2006-01-02"
	summary_format      = "worked for %.f:%2.f\n"
)

func (self *Command) Handle() {
	checks.InitCheck()
	time_now := time.Now()
	time_start := BeginingOfDay(time_now)
	time_end := EndOfDay(time_now)
	results, err := bench.EntriesBetween(&self.DB, time_start, time_end)
	if err != nil {
		fmt.Println("error getting entries")
		fmt.Println(err)
		os.Exit(1)
	}
	if results == nil {
		fmt.Println("nothing")
	} else {
		showEntries(results)
		duration := time_now.Sub(results[0].Time)
		fmt.Println("")
		fmt.Printf(summary_format, math.Floor(duration.Hours()), math.Floor(math.Mod(duration.Minutes(), 60)))
	}
}

func (self Command) Usage() {
	usage()
}
func (Command) Names() []string {
	return []string{"status"}
}
