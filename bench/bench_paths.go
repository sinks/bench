package bench

import (
	"os"
)

func BenchPath() string {
	return "./.bench"
}

func DbPath() string {
	return "./.bench/bench.db"
}

func BenchDirExists() bool {
	_, err := os.Stat(BenchPath())
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func DbExists() bool {
	_, err := os.Stat(DbPath())
	if os.IsNotExist(err) {
		return false
	}
	return true
}
