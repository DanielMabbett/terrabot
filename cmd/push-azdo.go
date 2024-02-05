package cmd

import (
	"context"
	"log"
	"terrabot/internal/azuredevops"
	"terrabot/internal/schema"

	azdo "github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/git"
	"github.com/spf13/cobra"
)

// Initialize the pushAzureDevOpsCmd and define its flags.
func init() {
	// Define flags for various parameters needed to push a comment to Azure DevOps.
	pushAzureDevOpsCmd.Flags().StringVarP(&organisation, "organisation", "o", "", "The organisation name to target in Azure DevOps Services.")
	pushAzureDevOpsCmd.Flags().StringVarP(&project, "project", "p", "", "The project name to target in Azure DevOps Services.")
	pushAzureDevOpsCmd.Flags().StringVarP(&repo, "repo", "r", "", "The Git repository name to target in Azure DevOps Services.")
	pushAzureDevOpsCmd.Flags().IntVarP(&pullRequestID, "pull-request-id", "i", 0, "The pull request ID for the Git Repo in Azure DevOps Services.")
	pushAzureDevOpsCmd.Flags().StringVarP(&personalAccessToken, "token", "t", "", "Your Personal Access Token (PAT) to Azure DevOps Services.")
	pushAzureDevOpsCmd.Flags().StringVarP(&planFile, "plan", "", "plan.txt", "The terraform plan file. Currently only supports .txt file outputs.")
	pushAzureDevOpsCmd.Flags().BoolVarP(&gruntBool, "grunt", "g", false, "Enable Terragrunt usage.")

	// Mark the required flags.
	pushAzureDevOpsCmd.MarkFlagRequired("organisation")
	pushAzureDevOpsCmd.MarkFlagRequired("project")
	pushAzureDevOpsCmd.MarkFlagRequired("repo")
	pushAzureDevOpsCmd.MarkFlagRequired("pull-request-id")
	pushAzureDevOpsCmd.MarkFlagRequired("personal-access-token")
	pushAzureDevOpsCmd.MarkFlagRequired("plan")
}

// Variable declarations for command flags.
var (
	organisation        string
	project             string
	repo                string
	pullRequestID       int
	personalAccessToken string
	planFile            string
	gruntBool           bool
)

// pushAzureDevOpsCmd represents the command to push comments to Azure DevOps pull requests.
var pushAzureDevOpsCmd = &cobra.Command{
	Use:   "azure-devops",
	Short: "Push a comment to a Pull Request to Azure DevOps services.",
	Run:   PushCommentToAzureDevOps,
}

// PushCommentToAzureDevOps pushes a comment to an Azure DevOps pull request.
func PushCommentToAzureDevOps(cmd *cobra.Command, args []string) {
	// Setup Azure DevOps connection using provided flags.
	connection := schema.AzureDevOpsConnection{
		Organisation:        organisation,
		Project:             project,
		Repository:          repo,
		PullRequestId:       pullRequestID,
		PersonalAccessToken: personalAccessToken,
	}

	// Construct the organization URL and establish a connection.
	organizationUrl := "https://dev.azure.com/" + connection.Organisation
	c := azdo.NewPatConnection(organizationUrl, connection.PersonalAccessToken)

	// Create a new context.
	ctx := context.Background()

	// Create a new Azure DevOps Git client.
	gitClient, err := git.NewClient(ctx, c)
	if err != nil {
		log.Printf("Error creating Azure DevOps Git client: %v\n", err)
		return
	}

	// Push comments based on the gruntBool flag.
	if gruntBool {
		azuredevops.PushTerragruntCommentToAzureDevOps(ctx, gitClient, connection.PullRequestId, connection.Project, connection.Repository, planFile)
	} else {
		azuredevops.PushTerraformCommentToAzureDevOps(ctx, gitClient, connection.PullRequestId, connection.Project, connection.Repository, planFile)
	}
}
