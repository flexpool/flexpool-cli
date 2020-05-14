package config

import "github.com/spf13/cobra"

// Init initializes the CLI command
func Init() {
	ConfigCmd.AddCommand(PathCmd)
	ConfigCmd.AddCommand(CleanCmd)
}

// ConfigCmd is the `flexpool-cli config` command
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Configuration management",
}
