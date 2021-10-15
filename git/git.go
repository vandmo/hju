package git

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SymbolicRef(folder string, symbol string) (string, error) {
	return runC(folder, "symbolic-ref", symbol)
}

type Summary struct {
	tracked   int
	untracked int
}

func Status(folder string) (*Summary, error) {
	lines, err := runC(folder, "status", "--porcelain")
	if err != nil {
		return nil, err
	}
	summary := Summary{}
	sc := bufio.NewScanner(strings.NewReader(lines))
	for sc.Scan() {
		line := sc.Text()
		if line[0] == '?' {
			summary.untracked += 1
		} else {
			summary.tracked++
		}
	}
	return &summary, nil
}

func HasCommit(folder string, commit string) (bool, error) {
	return runC_isSuccess(folder, "rev-parse", "--quiet", "--verify", commit+"^{commit}")
}

func HasRef(folder string, ref string) (bool, error) {
	return runC_isSuccess(folder, "show-ref", "--quiet", ref)
}

func FastForward(folder string) error {
	fmt.Println("--- \033[32mFast forwarding " + folder + "\033[0m")
	return run("-C", folder, "pull", "--ff-only")
}

func Fetch(folder string) error {
	fmt.Println("--- \033[32mFetching " + folder + "\033[0m")
	return run("-C", folder, "fetch")
}

func Switch(folder string, branch string, create bool) error {
	hasBranch, err := HasRef(folder, branch)
	if err != nil {
		return err
	}

	if hasBranch {
		fmt.Printf("--- \033[32mSwitching to branch %s in %s\033[0m\n", branch, folder)
		return run("-C", folder, "switch", branch)
	} else if create {
		fmt.Printf("--- \033[32mCreating and switching to branch %s in %s\033[0m\n", branch, folder)
		return run("-C", folder, "switch", "-c", branch)
	} else {
		fmt.Printf("--- \033[32mNOT creating NOR switching to branch %s in %s\033[0m\n", branch, folder)
		return nil
	}
}

func PrintDivergence(folder string, commit string) error {
	hasCommit, hasCommitErr := HasCommit(folder, commit)
	if hasCommitErr != nil {
		return hasCommitErr
	}
	if !hasCommit {
		fmt.Printf("%s \033[31m[no-such-commit]\033[0m\n", folder)
		return nil
	}
	lines, revListErr := runC(folder, "rev-list", "--left-right", "--count", "HEAD..."+commit)
	if revListErr != nil {
		return revListErr
	}
	fields := strings.Fields(lines)
	if fields[0] == "0" && fields[1] == "0" {
		fmt.Printf("%s (\033[36mup-to-date\033[0m)", folder)
	} else {
		fmt.Printf("%s \033[33m[A%s,B%s]\033[0m", folder, fields[0], fields[1])
	}
	fmt.Println()
	return nil
}

func PrintStatus(folder string) error {
	ref, err := SymbolicRef(folder, "HEAD")
	if err != nil {
		return err
	}
	lastSlashInd := strings.LastIndex(ref, "/")
	branch := strings.TrimSpace(ref[lastSlashInd+1:])
	status, statusErr := Status(folder)
	if statusErr != nil {
		return statusErr
	}
	fmt.Printf("\033[32m%s\033[0m (\033[36m%s\033[0m)", folder, branch)
	if status.tracked > 0 || status.untracked > 0 {
		fmt.Printf(" \033[31m[T%d,U%d]\033[0m", status.tracked, status.untracked)
	}
	fmt.Println()
	return nil
}

func Clone(url string) error {
	folderName := FolderName(url)
	if _, err := os.Stat(folderName); err == nil {
		fmt.Println("--- \033[33mNOT cloning: " + url + "\033[0m")
		return nil
	}
	fmt.Println("--- \033[32mCloning: " + url + "\033[0m")
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
