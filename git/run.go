package git

import (
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
