package git

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/cli/safeexec"
)

func command(arg ...string) (*exec.Cmd, error) {
	gitBin, err := safeexec.LookPath("git")
	if err != nil {
		return nil, err
	}
	return exec.Command(gitBin, arg...), nil
}

func run(arg ...string) error {
	cmd, err := command(arg...)
	if err != nil {
		return err
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func runC(folder string, arg0 string, arg1 string) (string, error) {
	cmd, err := command("-C", folder, arg0, arg1)
	if err != nil {
		return "", err
	}
	out, cmdErr := cmd.Output()
	if cmdErr != nil {
		return "", cmdErr
	}
	return string(out), nil
}

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

func HasBranch(folder string, branch string) (bool, error) {
	cmd, err := command("-C", folder, "rev-parse", "--verify", "--quiet", branch)
	if err != nil {
		return false, err
	}
	return cmd.Run() == nil, nil
}

func FastForward(folder string) error {
	fmt.Println("--- \033[32mFᴀsᴛ Fᴏʀᴡᴀʀᴅɪɴɢ " + folder + "\033[0m")
	return run("-C", folder, "pull", "--ff-only")
}

func Switch(folder string, branch string, create bool) error {
	hasBranch, err := HasBranch(folder, branch)
	if err != nil {
		return err
	}

	if hasBranch {
		fmt.Printf("--- \033[32mSwitching to branch %s in %s\033[0m\n", branch, folder)
		return run("-C", folder, "switch", branch)
	} else {
		fmt.Printf("--- \033[32mCreating and switching to branch %s in %s\033[0m\n", branch, folder)
		return run("-C", folder, "switch", "-c", branch)
	}
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
		fmt.Println("--- \033[33mℕ𝕆𝕋 𝕔𝕝𝕠𝕟𝕚𝕟𝕘: " + url + "\033[0m")
		return nil
	}
	fmt.Println("--- \033[32mℂ𝕃𝕆ℕ𝕀ℕ𝔾: " + url + "\033[0m")
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
