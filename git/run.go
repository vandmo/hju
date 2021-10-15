package git

import (
	"errors"
	"os"
	"os/exec"

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

func runC(folder string, arg ...string) (string, error) {
	args := make([]string, 0, 2+len(arg))
	args = append(args, "-C", folder)
	args = append(args, arg...)
	cmd, err := command(args...)
	if err != nil {
		return "", err
	}
	out, cmdErr := cmd.Output()
	if cmdErr != nil {
		return "", cmdErr
	}
	return string(out), nil
}

func runC_isSuccess(folder string, arg ...string) (bool, error) {
	args := make([]string, 0, 2+len(arg))
	args = append(args, "-C", folder)
	args = append(args, arg...)
	cmd, err := command(args...)
	if err != nil {
		return false, err
	}
	runErr := cmd.Run()
	if runErr == nil {
		return true, nil
	}
	var exitError *exec.ExitError
	if errors.As(runErr, &exitError) && exitError.ExitCode() == 1 {
		return false, nil
	}
	return false, runErr
}
