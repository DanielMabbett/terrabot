package cmd

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/git"
	"github.com/spf13/cobra"
)

func init() {
	// rootCmd.AddCommand(versionCmd)

	pushAzureDevOpsCmd.Flags().StringVarP(&organisation, "organisation", "o", "", "the organisation name to target in Azure DevOps Services.")
	pushAzureDevOpsCmd.Flags().StringVarP(&project, "project", "p", "", "the project name to target in Azure DevOps Services.")
	pushAzureDevOpsCmd.Flags().StringVarP(&repo, "repo", "r", "", "the git repository name to target in Azure DevOps Services.")
	pushAzureDevOpsCmd.Flags().IntVarP(&pullRequestID, "pull-request-id", "i", 0, "The pull request ID in for the Git Repo in Azure DevOps Services.")
	pushAzureDevOpsCmd.Flags().StringVarP(&personalAccessToken, "token", "t", "", "Your PAC to Azure DevOps Services.")
	pushAzureDevOpsCmd.Flags().StringVarP(&planFile, "plan", "", "plan.txt", "The terraform plan file. Currently only supports .txt file outputs.")
	pushAzureDevOpsCmd.MarkFlagRequired("organisation")
	pushAzureDevOpsCmd.MarkFlagRequired("project")
	pushAzureDevOpsCmd.MarkFlagRequired("repo")
	pushAzureDevOpsCmd.MarkFlagRequired("pull-request-id")
	pushAzureDevOpsCmd.MarkFlagRequired("personal-access-token")
	pushAzureDevOpsCmd.MarkFlagRequired("plan")
}

var organisation string
var project string
var repo string
var pullRequestID int
var personalAccessToken string
var planFile string

var pushAzureDevOpsCmd = &cobra.Command{
	Use:   "azure-devops",
	Short: "Push a comment to a Pull Request to Azure DevOps services.",
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

		b, err := ioutil.ReadFile(planFile)
		if err != nil {
			panic(err)
		}
		s := string(b)

		f, err := os.Open(planFile)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)

		line := 1

		plan := ""
		for scanner.Scan() {
			if strings.Contains(scanner.Text(), "Plan:") {
				plan = scanner.Text()
			}

			line++
		}
		if err := scanner.Err(); err != nil {
			panic(err)
		}

		fmt.Println(plan)
		overview := plan
		overviewWrap := "`" + overview + "`"

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
