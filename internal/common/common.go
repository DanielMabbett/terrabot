package common

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// RetrieveFileContents - Fetches the file contents of a plan file
func RetrieveFileContents(planFilePath string) (overview string, details string) {

	b, err := os.ReadFile(planFilePath)
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

// GenerateMarkdownContent generates markdown formatted content from a slice of strings and a fullDetails string.
func GenerateMarkdownContent(directories []string, fullDetails string) string {
	// Create a builder for efficient string concatenation
	var sb strings.Builder

	// Add the header and the overview section
	sb.WriteString("# Terrabot Response\n\n")
	sb.WriteString(fmt.Sprintf("### Overview\nThere are %d directories to apply\n\n", len(directories)))

	// Add the directories section
	sb.WriteString("### Directories\n")
	for _, dir := range directories {
		sb.WriteString(fmt.Sprintf("`%s`\n", dir)) // Each directory in code format
	}

	// Add the full output section
	sb.WriteString("\n### Full Output\n")
	sb.WriteString("<details><summary>Show Output</summary>\n\n```\n") // Use details tag for collapsible section
	sb.WriteString(fullDetails)
	sb.WriteString("\n```\n</details>\n")

	return sb.String()
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

// PrintStrings prints each element of a []string on a new line.
func PrintStrings(strings []string) {
	for _, str := range strings {
		fmt.Println(str)
	}
}
