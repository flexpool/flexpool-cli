package stat

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/ethereum/go-ethereum/common"
	"github.com/flexpool/flexpool-cli/address"
	"github.com/flexpool/flexpool-cli/api"
	"github.com/flexpool/flexpool-cli/utils"
	"github.com/spf13/cobra"
)

var charts bool

// Init initializes the CLI command
func Init() {
	StatCmd.Flags().BoolVarP(&charts, "charts", "c", false, "Show charts")
}

// StatCmd is the `flexpool-cli stat` command
var StatCmd = &cobra.Command{
	Use:   "stat [ADDRESS/NONE]",
	Short: "View the address' stats",
	Args:  cobra.MaximumNArgs(1),
	Run:   statCmd,
}

func statCmd(cmd *cobra.Command, args []string) {
	var addr common.Address
	if len(args) > 0 {
		if !common.IsHexAddress(args[0]) {
			fmt.Println("stat: invalid address")
			os.Exit(1)
		}
		addr = common.HexToAddress(args[0])
	} else {
		addresses := address.GetAddresses()
		if len(addresses) == 0 {
			fmt.Println("stat: No addresses configured. Use flexpool-cli [ADDRESS] or configure address watchlist.")
			os.Exit(1)
		} else if len(addresses) == 1 {
			addr = common.HexToAddress(addresses[0])
		} else {
			for i, address := range addresses {
				fmt.Println(" " + strconv.Itoa(i) + ") " + address)
			}

			reader := bufio.NewReader(os.Stdin)
			fmt.Print("\033[1;97mSelect address: \033[1;0m")
			addrIndexStr, _ := reader.ReadString('\n')
			addrIndex, err := strconv.Atoi(addrIndexStr[:len(addrIndexStr)-1])
			if err != nil {
				fmt.Println("stat: invalid address index")
				os.Exit(1)
			}
			if addrIndex >= len(addresses)+1 {
				fmt.Println("stat: address index out of range")
				os.Exit(1)
			}
			addr = common.HexToAddress(addresses[addrIndex])
		}
	}

	printStats(addr)
}

func printStats(address common.Address) {
	fmt.Println("\033[H\033[2J\033[1;32m- \033[1;97m" + address.String() + "\033[0m")

	minerStats, err := api.MinerStats(address.String())
	if err != nil {
		fmt.Println("\033[1;31mERROR:", err.Error())
		os.Exit(1)
	}

	fmt.Println("\n\033[1;97mStats:")
	workerCount, err := api.MinerWorkerCount(address.String())
	fmt.Print("    \033[1;97mWorkers: ")
	if err == nil {
		fmt.Println("\033[1;32m" + strconv.Itoa(workerCount.Online) + "\033[1;97m/\033[1;31m" + strconv.Itoa(workerCount.Offline))
	} else {
		fmt.Println("\033[1;31mERROR:", err.Error())
	}

	balance, err := api.MinerBalance(address.String())
	fmt.Print("    \033[1;97mBalance: ")
	var balanceFmt string
	if balance.Cmp(big.NewInt(0)) == 0 {
		balanceFmt = "0"
	} else {
		balanceFmt = fmt.Sprintf("%.5f", utils.WeiToEther(balance))
	}
	if err == nil {
		fmt.Println("\033[1;32m" + balanceFmt + " ETH")
	} else {
		fmt.Println("\033[1;31mERROR:", err.Error())
	}

	totalShares := minerStats.Daily.ValidShares + minerStats.Daily.StaleShares + minerStats.Daily.InvalidShares

	fmt.Print(" \033[1;97mEfficiency: ")
	efficiency := float64(minerStats.Daily.ValidShares) / float64(totalShares) * 100
	var efficiencyFmt string
	if math.IsNaN(efficiency) {
		efficiencyFmt = "0%"
	} else {
		efficiencyFmt = fmt.Sprintf("%.1f", efficiency) + "%"
	}
	fmt.Println("\033[1;32m" + efficiencyFmt)

	fmt.Println("\n\033[1;97mHashrates:")
	fmt.Println(" \033[1;97mEffective: \033[1;32m" + humanize.SIWithDigits(minerStats.Current.EffectiveHashrate, 1, "H/s") + "\033[1;97m (" + humanize.SIWithDigits(minerStats.Daily.EffectiveHashrate, 1, "H/s") + ")")
	fmt.Println("  \033[1;97mReported: \033[1;32m" + humanize.SIWithDigits(minerStats.Current.ReportedHashrate, 1, "H/s") + "\033[1;97m (" + humanize.SIWithDigits(minerStats.Current.ReportedHashrate, 1, "H/s") + ")")

	fmt.Println("\n\033[1;97mShares:")
	fmt.Println("   \033[1;97mValid: \033[1;32m" + strconv.FormatUint(minerStats.Daily.ValidShares, 10))
	fmt.Println("   \033[1;97mStale: \033[1;32m" + strconv.FormatUint(minerStats.Daily.StaleShares, 10))
	fmt.Println(" \033[1;97mInvalid: \033[1;32m" + strconv.FormatUint(minerStats.Daily.InvalidShares, 10))

	workers, _ := api.MinerWorkers(address.String())
	fmt.Println("\n\033[1;97mWorkers:")
	var workerNames []string
	var workerEffectiveFmt = make(map[string]string)
	var workerReportedFmt = make(map[string]string)
	var workerEfficiencyFmt = make(map[string]string)
	var workerValidFmt = make(map[string]string)
	var workerStaleFmt = make(map[string]string)
	var workerInvalidFmt = make(map[string]string)
	for _, worker := range workers {
		workerNames = append(workerNames, worker.Name)
		workerEffectiveFmt[worker.Name] = humanize.SIWithDigits(worker.EffectiveHashrate, 1, "H/s")
		workerReportedFmt[worker.Name] = humanize.SIWithDigits(worker.ReportedHashrate, 1, "H/s")
		workerEfficiencyFmt[worker.Name] = fmt.Sprintf("%.1f", float64(worker.ValidShares)/float64(worker.ValidShares+worker.StaleShares+worker.InvalidShares)*100) + "%"
		workerValidFmt[worker.Name] = strconv.FormatUint(worker.ValidShares, 10)
		workerStaleFmt[worker.Name] = strconv.FormatUint(worker.StaleShares, 10)
		workerInvalidFmt[worker.Name] = strconv.FormatUint(worker.InvalidShares, 10)
	}

	var longestEffectiveFmt int
	var longestReportedFmt int
	var longestEfficiencyFmt = 5
	var longestValidFmt int
	var longestStaleFmt int
	var longestInvalidFmt int

	for _, workerName := range workerNames {
		var length int
		length = len(workerEffectiveFmt[workerName])
		if length > longestEffectiveFmt {
			longestEffectiveFmt = length
		}
		length = len(workerReportedFmt[workerName])
		if length > longestReportedFmt {
			longestReportedFmt = length
		}
		length = len(workerValidFmt[workerName])
		if length > longestValidFmt {
			longestValidFmt = length
		}
		length = len(workerStaleFmt[workerName])
		if length > longestStaleFmt {
			longestStaleFmt = length
		}
		length = len(workerInvalidFmt[workerName])
		if length > longestInvalidFmt {
			longestInvalidFmt = length
		}
	}

	for _, workerName := range workerNames {
		fmt.Println(
			"  \033[1;97m" + workerName +
				" \033[0;37mR \033[1;32m" + workerReportedFmt[workerName] + strings.Repeat(" ", longestReportedFmt-len(workerReportedFmt[workerName])) +
				" \033[0;37mE \033[1;32m" + workerEffectiveFmt[workerName] + strings.Repeat(" ", longestEffectiveFmt-len(workerEffectiveFmt[workerName])) +
				" \033[0;37mEF \033[1;36m" + workerEfficiencyFmt[workerName] + strings.Repeat(" ", longestEfficiencyFmt-len(workerEfficiencyFmt[workerName])) +
				" \033[0;37mV \033[1;36m" + workerValidFmt[workerName] + strings.Repeat(" ", longestValidFmt-len(workerValidFmt[workerName])) +
				" \033[0;37mS \033[1;36m" + workerStaleFmt[workerName] + strings.Repeat(" ", longestStaleFmt-len(workerStaleFmt[workerName])) +
				" \033[0;37mI \033[1;36m" + workerInvalidFmt[workerName] + strings.Repeat(" ", longestInvalidFmt-len(workerInvalidFmt[workerName])) +
				"\033[0m")
	}

	if charts {
		chartData, _ := api.MinerChart(address.String())
		var effectiveHashrateChart []float64
		var reportedHashrateChart []float64
		for _, dat := range chartData {
			effectiveHashrateChart = append([]float64{dat.EffectiveHashrate}, effectiveHashrateChart...)
			reportedHashrateChart = append([]float64{dat.ReportedHashrate}, reportedHashrateChart...)
		}
		fmt.Println("\n\033[1;32m")
		fmt.Println(utils.Chart(effectiveHashrateChart, "Effective Hashrate", "H/s"))
		fmt.Println()
		fmt.Println(utils.Chart(reportedHashrateChart, "Reported Hashrate", "H/s"))
	}
	fmt.Println("\033[0m")
}
