package cmd

import (
	"context"
	"fmt"

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
	// todo: Consider making this a persistent flag
	pushAzureDevOpsCmd.Flags().BoolVarP(&gruntBool, "grunt", "g", false, "Enable terragrunt usage")
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
var gruntBool bool

var pushAzureDevOpsCmd = &cobra.Command{
	Use:   "azure-devops",
	Short: "Push a comment to a Pull Request to Azure DevOps services.",
	Run: func(cmd *cobra.Command, args []string) {

		organizationUrl := "https://dev.azure.com/" + organisation
		connection := azuredevops.NewPatConnection(organizationUrl, personalAccessToken)

		ctx := context.Background()

		gitClient, err := git.NewClient(ctx, connection)
		if err != nil {
			fmt.Println("error")
			fmt.Println(err)
			return
		}

		if gruntBool {
			fmt.Println("[information] terragrunt plans are partially supported (no overviews). Continuing...")

			details := RetrieveGruntFileContents(planFile)
			content := GenerateContentString("Terragrunt overview not supported", details)

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
			fmt.Println("Sucessfully pushed thread to Azure DevOps Services.")

		} else {

			overview, details := RetrieveFileContents(planFile)
			content := GenerateContentString(overview, details)

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
			fmt.Println("Sucessfully pushed thread to Azure DevOps Services.")
		}

	},
}
