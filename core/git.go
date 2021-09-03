package core

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"github.com/cli/safeexec"
)

func GitIt(arg ...string) error {
	gitBin, err := safeexec.LookPath("git")
	if err != nil {
		return err
	}
	cmd := exec.Command(gitBin, arg...)
	return cmd.Run()
}

type GitIgnore struct {
	lines []string
}

func ParseGitIgnore() (*GitIgnore, error) {
	file, err := os.Open(".gitignore")
	if err != nil {
		if os.IsNotExist(err) {
			return &GitIgnore{}, nil
		} else {
			return nil, err
		}
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	scannerErr := scanner.Err()
	if scannerErr != nil {
		return nil, scannerErr
	}
	return &GitIgnore{lines: lines}, nil
}

func (gi *GitIgnore) Write() error {
	if gi.lines == nil {
		return nil
	}
	file, err := os.Create(".gitignore")
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range gi.lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func (gi *GitIgnore) AddIfNeeded(line string) {
	if gi.lines == nil {
		gi.lines = make([]string, 0)
	}
	gi.lines = append(gi.lines, line)
}
