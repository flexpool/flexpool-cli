package address

import "github.com/spf13/cobra"

// Init initializes the CLI command
func Init() {
	AddressCmd.AddCommand(ListCmd)

	addInit()
	AddressCmd.AddCommand(AddCmd)

	AddressCmd.AddCommand(RemoveCmd)
}

// AddressCmd is the `flexpool-cli addr` command
var AddressCmd = &cobra.Command{
	Use:   "addr [OPTIONS]",
	Short: "Address management",
}
