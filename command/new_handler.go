package command

import (
	"fmt"
	"os"
)

const dir_perms = 0755

type NewCommandCreatePathError string

func (s NewCommandCreatePathError) Error() string {
	return fmt.Sprintf("failed to create %s", string(s))
}

func NewHandler() {
	err := newBenchDir(basePath())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println("bench created")
	}
}

func basePath() string {
	return ".bench/"
}

func newBenchDir(base_path string) error {
	err := os.MkdirAll(base_path, dir_perms)
	if err != nil {
		return NewCommandCreatePathError(base_path)
	}
	return nil
}
