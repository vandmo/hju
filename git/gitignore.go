package git

import (
	"bufio"
	"fmt"
	"os"
)

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

func (gi *GitIgnore) Contains(line string) bool {
	if gi.lines != nil {
		for _, existingLine := range gi.lines {
			if line == existingLine {
				return true
			}
		}
	}
	return false
}

func (gi *GitIgnore) AddRootFolderIfNeeded(line string) {
	gi.AddIfNeeded("/" + line + "/")
}

func (gi *GitIgnore) Remove(entry string) {
	old := gi.lines
	gi.lines = gi.lines[:0]
	for _, line := range old {
		if entry != line {
			gi.lines = append(gi.lines, line)
		}
	}
}

func (gi *GitIgnore) RemoveRootFolder(line string) {
	gi.Remove("/" + line + "/")
}

func (gi *GitIgnore) AddIfNeeded(line string) {
	if gi.lines == nil {
		gi.lines = make([]string, 0)
	}
	if gi.Contains(line) {
		return
	}
	gi.lines = append(gi.lines, line)
}
