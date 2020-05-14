package config

import (
	"fmt"
	"os"

	"github.com/flexpool/flexpool-cli/utils"
	"github.com/spf13/cobra"
)

// CleanCmd is the `flexpool-cli address list` command
var CleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Cleans the configuration path",
	Args:  cobra.NoArgs,
	Run:   cleanCmd,
}

func cleanCmd(cmd *cobra.Command, args []string) {
	path, exists := getPath()
	if !exists {
		fmt.Println("clean: no config path found")
		os.Exit(1)
	}
	fmt.Print("clean: remove " + path + "? ")
	if !utils.Ask4confirm() {
		fmt.Println("clean: cancelled")
		os.Exit(0)
	}
	err := os.RemoveAll(path)
	if err != nil {
		fmt.Println("clean: " + err.Error())
		os.Exit(1)
	}
	fmt.Println("Removed " + path)
}
