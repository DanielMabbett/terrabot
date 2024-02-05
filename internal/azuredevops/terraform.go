package azuredevops

import (
	"context"
	"fmt"
	"log" // Import the log package
	"terrabot/internal/common"

	"github.com/microsoft/azure-devops-go-api/azuredevops/git"
)

func PushTerraformCommentToAzureDevOps(ctx context.Context, gitClient git.Client, pullRequestID int, project string, repo string, planFile string) {
	// Log the start of the Terraform push to Azure DevOps.
	log.Println("Starting Terraform push to Azure DevOps...")

	// Retrieve the content from the plan file.
	overview, details := common.RetrieveFileContents(planFile)
	content := common.GenerateContentString(overview, details)

	// Create a comment thread with the content.
	thread := git.CreateThreadArgs{
		CommentThread: &git.GitPullRequestCommentThread{
			Comments: &[]git.Comment{
				{
					Content: common.PtrString(content),
				},
			},
		},
		PullRequestId: &pullRequestID,
		Project:       &project,
		RepositoryId:  &repo,
	}

	// Attempt to create the thread on the pull request.
	_, err := gitClient.CreateThread(ctx, thread)
	if err != nil {
		fmt.Println("Error on Creating Thread")
		fmt.Println(err)
		return
	}
	// Log successful push to Azure DevOps Services.
	fmt.Println("Successfully pushed thread to Azure DevOps Services.")
}
