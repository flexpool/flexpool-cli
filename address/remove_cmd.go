package address

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

// RemoveCmd is the `flexpool-cli addr rm` command
var RemoveCmd = &cobra.Command{
	Use:   "rm [ADDRESS]",
	Short: "Remove address from watchlist",
	Args:  cobra.MinimumNArgs(1),
	Run:   removeCmd,
}

func removeCmd(cmd *cobra.Command, args []string) {
	for _, addr := range args {
		if !common.IsHexAddress(addr) {
			fmt.Println("rm: invalid address")
			break
		}
		address := common.HexToAddress(addr).String()
		var deleted bool
		var newAddresses []string
		for _, addr := range GetAddresses() {
			if addr != address {
				newAddresses = append(newAddresses, addr)
			} else {
				deleted = true
			}
		}

		if deleted {
			setAddresses(newAddresses)
		} else {
			fmt.Println("rm: address does not exists")
			break
		}
		fmt.Println(address)
	}
}
