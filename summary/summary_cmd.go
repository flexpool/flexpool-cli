package summary

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/flexpool/flexpool-cli/address"
	"github.com/flexpool/flexpool-cli/api"
	"github.com/spf13/cobra"
)

var average bool

// Init initializes the CLI command
func Init() {
	SummaryCmd.Flags().BoolVarP(&average, "average", "a", false, "Show average values")
}

// SummaryCmd is the `flexpool-cli sum` command
var SummaryCmd = &cobra.Command{
	Use:   "sum",
	Short: "View the summary on watched addresses",
	Args:  cobra.NoArgs,
	Run:   sumCmd,
}

func sumCmd(cmd *cobra.Command, args []string) {
	addresses := address.GetAddresses()
	var sumReported float64
	var sumEffective float64
	var statsMap = make(map[string]api.DailyCurrentStats)
	var workersCountMap = make(map[string]string)
	var erroredCount int
	var notErroredAddresses []string

	for _, addr := range addresses {
		minerStats, err := api.MinerStats(addr)
		if err != nil {
			erroredCount++
			continue
		}
		workers, err := api.MinerWorkerCount(addr)
		if err != nil {
			erroredCount++
			continue
		}

		statsMap[addr] = minerStats
		workersCountMap[addr] = strconv.Itoa(workers.Online) + "/" + strconv.Itoa(workers.Offline)

		notErroredAddresses = append(notErroredAddresses, addr)
		if average {
			sumReported += minerStats.Daily.ReportedHashrate
			sumEffective += minerStats.Daily.EffectiveHashrate
		} else {
			sumReported += minerStats.Current.ReportedHashrate
			sumEffective += minerStats.Current.EffectiveHashrate
		}
	}

	var longestWorkerCountString int
	var longestReportedString int
	var longestEffectiveString int
	var longestEfficiencyString int
	var longestValidSharesString int
	var longestStaleSharesString int
	var longestInvalidSharesString int

	var workerCountFmtStringMap = make(map[string]string)
	var reportedFmtStringMap = make(map[string]string)
	var effectiveFmtStringMap = make(map[string]string)
	var efficiencyFmtMap = make(map[string]string)
	var validSharesFmtStringMap = make(map[string]string)
	var staleSharesFmtStringMap = make(map[string]string)
	var invalidSharesFmtStringMap = make(map[string]string)

	var length int

	for _, addr := range notErroredAddresses {
		workerCountFmtStringMap[addr] = workersCountMap[addr]
		length = len(workerCountFmtStringMap[addr])
		if length > longestWorkerCountString {
			longestWorkerCountString = length
		}
		var reported float64
		if average {
			reported = statsMap[addr].Daily.ReportedHashrate
		} else {
			reported = statsMap[addr].Current.ReportedHashrate
		}
		reportedFmtStringMap[addr] = humanize.SIWithDigits(reported, 1, "H/s")
		length = len(reportedFmtStringMap[addr])
		if length > longestReportedString {
			longestReportedString = length
		}
		var effective float64
		if average {
			effective = statsMap[addr].Daily.EffectiveHashrate
		} else {
			effective = statsMap[addr].Current.EffectiveHashrate
		}
		effectiveFmtStringMap[addr] = humanize.SIWithDigits(effective, 1, "H/s")
		length = len(reportedFmtStringMap[addr])
		if length > longestEffectiveString {
			longestEffectiveString = length
		}
		totalShares := statsMap[addr].Daily.ValidShares + statsMap[addr].Daily.StaleShares + statsMap[addr].Daily.InvalidShares
		if totalShares > 0 {
			efficiencyFmtMap[addr] = fmt.Sprintf("%.1f", float64(statsMap[addr].Daily.ValidShares)/float64(totalShares)*100) + "%"
		} else {
			efficiencyFmtMap[addr] = "0%"
		}
		length = len(efficiencyFmtMap[addr])
		if length > longestEfficiencyString {
			longestEfficiencyString = length
		}
		validSharesFmtStringMap[addr] = strconv.FormatUint(statsMap[addr].Daily.ValidShares, 10)
		length = len(validSharesFmtStringMap[addr])
		if length > longestValidSharesString {
			longestValidSharesString = length
		}
		staleSharesFmtStringMap[addr] = strconv.FormatUint(statsMap[addr].Daily.StaleShares, 10)
		length = len(staleSharesFmtStringMap[addr])
		if length > longestStaleSharesString {
			longestStaleSharesString = length
		}
		invalidSharesFmtStringMap[addr] = strconv.FormatUint(statsMap[addr].Daily.InvalidShares, 10)
		length = len(invalidSharesFmtStringMap[addr])
		if length > longestInvalidSharesString {
			longestInvalidSharesString = length
		}
	}

	fmt.Println()
	for _, addr := range notErroredAddresses {
		fmt.Println(
			"\033[1;97m" + addr[:7] + "…" + addr[40-3:] +
				" \033[0;37mW \033[1;36m" + workerCountFmtStringMap[addr] + strings.Repeat(" ", longestWorkerCountString-len(workerCountFmtStringMap[addr])) +
				" \033[0;37mR \033[1;32m" + reportedFmtStringMap[addr] + strings.Repeat(" ", longestReportedString-len(reportedFmtStringMap[addr])) +
				" \033[0;37mE \033[1;32m" + effectiveFmtStringMap[addr] + strings.Repeat(" ", longestEffectiveString-len(effectiveFmtStringMap[addr])) +
				" \033[0;37mEF \033[1;36m" + efficiencyFmtMap[addr] + strings.Repeat(" ", longestEfficiencyString-len(efficiencyFmtMap[addr])) +
				" \033[0;37mV \033[1;36m" + validSharesFmtStringMap[addr] + strings.Repeat(" ", longestValidSharesString-len(validSharesFmtStringMap[addr])) +
				" \033[0;37mS \033[1;36m" + staleSharesFmtStringMap[addr] + strings.Repeat(" ", longestStaleSharesString-len(staleSharesFmtStringMap[addr])) +
				" \033[0;37mI \033[1;36m" + invalidSharesFmtStringMap[addr] + strings.Repeat(" ", longestInvalidSharesString-len(invalidSharesFmtStringMap[addr])) +
				"\033[0m")
	}

	fmt.Println()
	fmt.Print(len(notErroredAddresses), " total ")
	if erroredCount > 0 {
		fmt.Print("(" + strconv.Itoa(erroredCount) + " unavailable) ")
	}
	fmt.Println("-", humanize.SIWithDigits(sumReported, 1, "H/s"), "(Effective "+humanize.SIWithDigits(sumEffective, 1, "H/s")+")")
}
