package cmd

import (
	"fmt"
    "os"
    "io/ioutil"
    "encoding/json"
	"github.com/spf13/cobra"
)

type Repositories struct {
    Repositories []string `json:"repositories"`
}

// cloneCmd represents the clone command
var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("clone called")
        // Open our jsonFile
        jsonFile, err := os.Open("hju.json")
        // if we os.Open returns an error then handle it
        if err != nil {
            fmt.Println(err)
        }
        fmt.Println("Successfully Opened hju.json")
        // defer the closing of our jsonFile so that we can parse it later on
        defer jsonFile.Close()

        // read our opened jsonFile as a byte array.
        byteValue, _ := ioutil.ReadAll(jsonFile)

        // we initialize our Users array
        var repositories Repositories

        // we unmarshal our byteArray which contains our
        // jsonFile's content into 'users' which we defined above
        json.Unmarshal(byteValue, &repositories)

        // we iterate through every user within our users array and
        // print out the user Type, their name, and their facebook url
        // as just an example
        for i := 0; i < len(repositories.Repositories); i++ {
            fmt.Println("Cloning: " + repositories.Repositories[i])
        }
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cloneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cloneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
