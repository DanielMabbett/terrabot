# terrabot
Push Terraform Plans back to your PRs and make your process more gitops!

The idea originally came from the https://github.com/runatlantis/atlantis tool so check this out first and see if it fulfills your requirements.

## Why use terrabot?

This is not designed to be a fully "fleshed-out" tool such as Atlantis. 
It focuses only on sending a comment back to your pull request to let the reviewers know details of the intended changes. 

## Usage

```sh
# Export the text of your terraform plan out
terraform plan -no-color > plan.txt

# Then run terrabot push azure-devops
# Note: For you to use this in an azure devops pipeline, make use of the pipeline variables 
# https://docs.microsoft.com/en-us/azure/devops/pipelines/build/variables?view=azure-devops&tabs=yaml
# set your PAC by export PAC=yourpac
terrabot push azure-devops --organisation myorg --project "my project" --repo myrepo --pull-request-id 1 --token $PAC
```

## Contributors 

Contributions are welcome! 

