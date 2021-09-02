package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/cli/safeexec"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

type Repositories struct {
	Repositories []string `json:"repositories"`
}

func git(arg ...string) error {
	gitBin, err := safeexec.LookPath("git")
	if err != nil {
		return err
	}
	cmd := exec.Command(gitBin, arg...)
	return cmd.Run()
}

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clones the needed repositories",
	RunE: func(cmd *cobra.Command, args []string) error {
		jsonFile, fileOpenErr := os.Open("hju.json")
		if fileOpenErr != nil {
			return fileOpenErr
		}
		defer jsonFile.Close()

		decoder := json.NewDecoder(jsonFile)
		var repositories Repositories
		decoderErr := decoder.Decode(&repositories)
		if decoderErr != nil {
			return decoderErr
		}

		for _, repo := range repositories.Repositories {
			fmt.Println("Cloning: " + repo)
			gitErr := git("clone", repo)
			if gitErr != nil {
				return gitErr
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)
}
