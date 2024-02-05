package grunt

import (
	"bufio"
	"os"
	"strings"
)

// RetrieveGruntFileContents fetches the lines containing "Plan:" from a grunt plan file and the entire content of the file.
func RetrieveGruntFileContents(planFilePath string) (details []string, fullContent string, err error) {
	// Open the file for reading.
	f, err := os.Open(planFilePath)
	if err != nil {
		return nil, "", err // Return the error to be handled by the caller.
	}
	defer f.Close()

	// Use a scanner to read the file line by line.
	scanner := bufio.NewScanner(f)
	var allLinesBuilder strings.Builder // Builder to accumulate all file contents.
	for scanner.Scan() {
		line := scanner.Text()
		allLinesBuilder.WriteString(line + "\n") // Add the line to the full content.
		// Check if the line contains "Plan:".
		if strings.Contains(line, "Plan:") {
			details = append(details, line) // Add the line to the results.
		}
	}

	// Check for any errors during scanning.
	if err := scanner.Err(); err != nil {
		return nil, "", err // Return the error to be handled by the caller.
	}

	// Get the full content from the builder.
	fullContent = allLinesBuilder.String()

	return details, fullContent, nil // Return the collected lines and the full content.
}
