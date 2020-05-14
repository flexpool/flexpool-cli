package address

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// ListCmd is the `flexpool-cli addr ls` command
var ListCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all watched addresses",
	Args:  cobra.NoArgs,
	Run:   listCmd,
}

func listCmd(cmd *cobra.Command, args []string) {
	addrs := GetAddresses()
	if len(addrs) > 0 {
		for _, addr := range addrs {
			fmt.Println("-", addr)
		}
		fmt.Println()
		fmt.Println(len(addrs), "total")
	} else {
		fmt.Println("ls: no addresses found")
		os.Exit(1)
	}

}
