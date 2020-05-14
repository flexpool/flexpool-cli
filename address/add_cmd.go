package address

import (
	"fmt"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/flexpool/flexpool-cli/api"
	"github.com/flexpool/flexpool-cli/utils"
	"github.com/spf13/cobra"
)

var force bool

func addInit() {
	AddCmd.Flags().BoolVarP(&force, "force", "f", false, "Skip all address checks")
}

// AddCmd is the `flexpool-cli addr add` command
var AddCmd = &cobra.Command{
	Use:   "add [ADDRESS]",
	Short: "Add a new address to watchlist",
	Args:  cobra.ExactArgs(1),
	Run:   addCmd,
}

func addCmd(cmd *cobra.Command, args []string) {
	if !common.IsHexAddress(args[0]) {
		fmt.Println("add: invalid address")
		os.Exit(1)
	}
	address := common.HexToAddress(args[0])
	if containsAddress(address.String()) {
		fmt.Println("add: duplicate address")
		os.Exit(1)
	}
	if !force {
		if args[0] != strings.ToLower(args[0]) && address.String() != args[0] {
			fmt.Print("add: checksum doesn't match, continue? ")
			if !utils.Ask4confirm() {
				fmt.Println("add: cancelled")
				os.Exit(0)
			}
		}
		exists, _ := api.MinerExists(address.String())
		if !exists {
			fmt.Print("add: address does not exist in flexpool database, continue? ")
			if !utils.Ask4confirm() {
				fmt.Println("add: cancelled")
				os.Exit(0)
			}
		}
	}
	setAddresses(append(GetAddresses(), address.String()))
	fmt.Println(address.String())
}
