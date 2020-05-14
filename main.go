package main

import (
	"github.com/flexpool/flexpool-cli/address"
	"github.com/flexpool/flexpool-cli/config"
	"github.com/flexpool/flexpool-cli/stat"
	"github.com/flexpool/flexpool-cli/summary"
	"github.com/flexpool/flexpool-cli/version"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "flexpool-cli"}

func main() {
	rootCmd.AddCommand(version.VersionCmd)

	config.Init()
	rootCmd.AddCommand(config.ConfigCmd)

	address.Init()
	rootCmd.AddCommand(address.AddressCmd)

	summary.Init()
	rootCmd.AddCommand(summary.SummaryCmd)

	stat.Init()
	rootCmd.AddCommand(stat.StatCmd)

	rootCmd.Execute()
}
