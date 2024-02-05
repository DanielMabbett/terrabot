package azuredevops

import (
	"context"
	"fmt"
	"log" // Import the log package
	"terrabot/internal/common"
	"terrabot/internal/grunt"

	"github.com/microsoft/azure-devops-go-api/azuredevops/git"
)

func PushTerragruntCommentToAzureDevOps(ctx context.Context, gitClient git.Client, pullRequestID int, project string, repo string, planFile string) {
	// Log the start of the Terragrunt push to Azure DevOps.
	log.Println("Starting Terragrunt push to Azure DevOps...")

	// Retrieve the Terragrunt file contents.
	details, fullContent, err := grunt.RetrieveGruntFileContents(planFile)
	if err != nil {
		fmt.Println("Error on Reading Terragrunt File Contents")
		fmt.Println(err)
		return
	}
	common.PrintStrings(details)

	content := common.GenerateMarkdownContent(details, fullContent)

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
	_, err = gitClient.CreateThread(ctx, thread)
	if err != nil {
		fmt.Println("Error on Creating Thread")
		fmt.Println(err)
		return
	}
	// Log successful push to Azure DevOps Services.
	fmt.Println("Successfully pushed thread to Azure DevOps Services.")
}
