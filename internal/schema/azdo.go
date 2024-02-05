package schema

type AzureDevOpsConnection struct {
	Organisation        string
	Project             string
	Repository          string
	PullRequestId       int
	PersonalAccessToken string
}
