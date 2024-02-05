# Makefile for a Go project.

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOFMT=$(GOCMD) fmt

# Binary names
BINARY_NAME=terrabot

all: test build
build: 
	$(GOBUILD) -o $(BINARY_NAME) -v
test: 
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
fmt:
	$(GOFMT) ./...
grunt:
	@bash ./scripts/terragrunt-plan.sh

.PHONY: all build test clean fmt grunt
