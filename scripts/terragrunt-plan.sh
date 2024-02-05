#!/bin/bash

# Define the target directory
TARGET_DIR="./test/grunt"

# Check if the target directory exists
if [ -d "$TARGET_DIR" ]; then
    # Navigate to the target directory
    cd "$TARGET_DIR" || exit  # Exit if changing directory fails
    
    # Run the terragrunt command
    terragrunt run-all plan -out tfplan
    if [ $? -ne 0 ]; then  # Check if terragrunt command was successful
        echo "Terragrunt command failed"
        exit 1
    fi
    
    terragrunt run-all show -json tfplan > tfplan.json

else
    echo "Target directory does not exist"
    exit 1
fi
