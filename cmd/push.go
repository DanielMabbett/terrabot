package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	// rootCmd.AddCommand(versionCmd)

	pushCmd.PersistentFlags().StringP("organisation", "o", "YOUR NAME", "the organisation name to target in Azure DevOps Services.")
	pushCmd.PersistentFlags().StringP("project", "p", "YOUR NAME", "the project name to target in Azure DevOps Services.")
	pushCmd.PersistentFlags().StringP("repo", "r", "YOUR NAME", "the git repository name to target in Azure DevOps Services.")
}

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push a comment to a Pull Request",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD" + repo)
	},
}
