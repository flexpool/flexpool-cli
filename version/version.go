package version

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

// Version is a flexpool-cli version
const Version = "v0.0.1"

// VersionCmd is the `flexpoolcli version` command
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Run:   versionCmd,
}

func versionCmd(cmd *cobra.Command, args []string) {
	fmt.Println("flexpool-cli", Version, "(Built on "+runtime.GOOS+"/"+runtime.GOARCH+" using "+runtime.Compiler+", "+runtime.Version()+")")
}
