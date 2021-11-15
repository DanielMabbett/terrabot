package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/git"
	"github.com/spf13/cobra"
)

func init() {
	// rootCmd.AddCommand(versionCmd)

	pushCmd.Flags().StringVarP(&organisation, "organisation", "o", "", "the organisation name to target in Azure DevOps Services.")
	pushCmd.Flags().StringVarP(&project, "project", "p", "", "the project name to target in Azure DevOps Services.")
	pushCmd.Flags().StringVarP(&repo, "repo", "r", "", "the git repository name to target in Azure DevOps Services.")
	pushCmd.Flags().IntVarP(&pullRequestID, "pull-request-id", "i", 0, "The pull request ID in for the Git Repo in Azure DevOps Services.")
	pushCmd.Flags().StringVarP(&personalAccessToken, "token", "t", "", "Your PAC to Azure DevOps Services.")
	pushCmd.Flags().StringVarP(&planFile, "plan", "", "plan.txt", "The terraform plan file. Currently only supports .txt file outputs.")

	pushCmd.MarkFlagRequired("organisation")
	pushCmd.MarkFlagRequired("project")
	pushCmd.MarkFlagRequired("repo")
	pushCmd.MarkFlagRequired("pull-request-id")
	pushCmd.MarkFlagRequired("personal-access-token")
	pushCmd.MarkFlagRequired("plan")
}

var organisation string
var project string
var repo string
var pullRequestID int
var personalAccessToken string
var planFile string

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push a comment to a Pull Request",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		// azdo alternative
		organizationUrl := "https://dev.azure.com/" + organisation
		connection := azuredevops.NewPatConnection(organizationUrl, personalAccessToken)

		ctx := context.Background()

		gitClient, err := git.NewClient(ctx, connection)
		if err != nil {
			fmt.Println("error")
			fmt.Println(err)
			return
		}

		var text string = "Plan:"

		// read the whole file at once
		b, err := ioutil.ReadFile(planFile)
		if err != nil {
			panic(err)
		}
		s := string(b)
		//check whether s contains substring text
		fmt.Println(strings.Contains(s, text))
		fmt.Println(strings.Trim(s, text))

		overviewWrap := "`overview. Comming soon.`"

		fullDetails := s
		fullDetailsWrap := "```\n" + fullDetails + "\n```"

		content := `🤖 Terrabot Response ⚡
						
Overview
` + overviewWrap + `
<details><summary>Full Details</summary>

` + fullDetailsWrap

		thread := git.CreateThreadArgs{
			CommentThread: &git.GitPullRequestCommentThread{
				Comments: &[]git.Comment{
					{
						Content: PtrString(content),
					},
				},
			},
			PullRequestId: &pullRequestID,
			Project:       &project,
			RepositoryId:  &repo,
		}

		_, err = gitClient.CreateThread(ctx, thread)
		if err != nil {
			fmt.Println("Error on Creating Thread")
			fmt.Println(err)
			return
		}

	},
}

func PtrString(v string) *string { return &v }
