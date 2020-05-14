package config

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// PathCmd is the `flexpool-cli address list` command
var PathCmd = &cobra.Command{
	Use:   "path",
	Short: "Show the configuration path",
	Args:  cobra.NoArgs,
	Run:   pathCmd,
}

func pathCmd(cmd *cobra.Command, args []string) {
	path, exists := getPath()
	if !exists {
		fmt.Println("path: no config path found")
		os.Exit(1)
	}
	fmt.Println(path)
}
