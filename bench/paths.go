package bench

import (
	"os"
)

const (
	BenchDir    = "./.bench/"
	BenchDBPath = BenchDir + "bench.db"
)

func BenchDirExists() bool {
	_, err := os.Stat(BenchDir)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func DbExists() bool {
	_, err := os.Stat(BenchDBPath)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
