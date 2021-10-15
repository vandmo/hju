package git

import (
	"bufio"
	"strconv"
	"strings"
)

func SymbolicRef(folder string, symbol string) (string, error) {
	return runC(folder, "symbolic-ref", symbol)
}

type Status struct {
	Tracked   int
	Untracked int
}

type Divergence struct {
	Ahead  int
	Behind int
}

func GetStatus(folder string) (*Status, error) {
	lines, err := runC(folder, "status", "--porcelain")
	if err != nil {
		return nil, err
	}
	status := Status{}
	sc := bufio.NewScanner(strings.NewReader(lines))
	for sc.Scan() {
		line := sc.Text()
		if line[0] == '?' {
			status.Untracked += 1
		} else {
			status.Tracked++
		}
	}
	return &status, nil
}

func HasCommit(folder string, commit string) (bool, error) {
	return runC_isSuccess(folder, "rev-parse", "--quiet", "--verify", commit+"^{commit}")
}

func HasRef(folder string, ref string) (bool, error) {
	return runC_isSuccess(folder, "show-ref", "--quiet", ref)
}

func FastForward(folder string) error {
	return run("-C", folder, "pull", "--ff-only")
}

func Fetch(folder string) error {
	return run("-C", folder, "fetch")
}

func Switch(folder string, branch string, create bool) error {
	if create {
		return run("-C", folder, "switch", "-c", branch)
	} else {
		return run("-C", folder, "switch", branch)
	}
}

func GetDivergence(folder string, commit string) (*Divergence, error) {
	lines, revListErr := runC(folder, "rev-list", "--left-right", "--count", "HEAD..."+commit)
	if revListErr != nil {
		return nil, revListErr
	}
	fields := strings.Fields(lines)
	ahead, parseErr := strconv.Atoi(fields[0])
	if parseErr != nil {
		return nil, parseErr
	}
	behind, parseErr := strconv.Atoi(fields[1])
	if parseErr != nil {
		return nil, parseErr
	}
	return &Divergence{Ahead: ahead, Behind: behind}, nil
}

func Clone(url string) error {
	return run("clone", url)
}

func FolderName(gitURL string) string {
	lastSlashInd := strings.LastIndex(gitURL, "/")
	basename := gitURL[lastSlashInd+1:]
	lastDotInd := strings.LastIndex(basename, ".")
	if lastDotInd > -1 {
		return basename[:lastDotInd]
	}
	return basename
}
