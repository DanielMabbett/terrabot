package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Terrabot",
	Long:  `All software has versions. This is Terrabot's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Terrabot Version -- v0.2")
	},
}
