package cmd

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

// RetrieveFileContents - Fetches the file contents of a plan file
func RetrieveFileContents(planFilePath string) (overview string, details string) {

	b, err := ioutil.ReadFile(planFilePath)
	if err != nil {
		panic(err)
	}
	s := string(b)

	f, err := os.Open(planFilePath)
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

	return plan, s
}

// RetrieveFileContents - Fetches the file contents of a grunt plan file
func RetrieveGruntFileContents(planFilePath string) (details string) {

	b, err := ioutil.ReadFile(planFilePath)
	if err != nil {
		panic(err)
	}
	s := string(b)

	f, err := os.Open(planFilePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	return s
}

// GenerateContentString - Generates the content string to be parsed into a request
func GenerateContentString(overview string, details string) string {
	overviewWrap := "`" + overview + "`"

	fullDetailsWrap := "```\n" + details + "\n```"

	content := `🤖 Terrabot Response ⚡
					
Overview
` + overviewWrap + `
<details><summary>Full Details</summary>

` + fullDetailsWrap

	return content
}

// PtrString - Convert string to pointer string
func PtrString(v string) *string { return &v }
