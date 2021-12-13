package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

// VersionCommand get bopher version
var VersionCommand = &cobra.Command{
	Use:   "version",
	Short: "get bopher version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v1.1.0")
	},
}
