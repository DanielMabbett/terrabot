package grunt

import (
	"os"
	"strings"
	"testing"
)

// createTestFile creates a file for testing purposes and returns its path and a cleanup function.
func createTestFile(content string, t *testing.T) (string, func()) {
	t.Helper()
	file, err := os.CreateTemp("", "grunt-test-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}

	if err := file.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	return file.Name(), func() { os.Remove(file.Name()) } // cleanup function
}

func TestRetrieveGruntFileContents(t *testing.T) {
	// Test case with expected "Plan:" lines
	t.Run("WithPlanLines", func(t *testing.T) {
		content := `Line1
Plan: A plan
Line2
Plan: Another plan
Line3`
		filePath, cleanup := createTestFile(content, t)
		defer cleanup()

		details, fullContent, err := RetrieveGruntFileContents(filePath)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if len(details) != 2 {
			t.Errorf("Expected 2 details, got %d", len(details))
		}
		if !strings.Contains(fullContent, "Line1") {
			t.Errorf("Expected full content to contain 'Line1', got: %s", fullContent)
		}
	})

	// Test case with no "Plan:" lines
	t.Run("WithoutPlanLines", func(t *testing.T) {
		content := `Line1
Line2
Line3`
		filePath, cleanup := createTestFile(content, t)
		defer cleanup()

		details, fullContent, err := RetrieveGruntFileContents(filePath)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if len(details) != 0 {
			t.Errorf("Expected 0 details, got %d", len(details))
		}
		if !strings.Contains(fullContent, "Line1") {
			t.Errorf("Expected full content to contain 'Line1', got: %s", fullContent)
		}
	})

	// Test case with invalid file path
	t.Run("InvalidFilePath", func(t *testing.T) {
		filePath := "/invalid/path"
		_, _, err := RetrieveGruntFileContents(filePath)
		if err == nil {
			t.Error("Expected error for invalid file path, got nil")
		}
	})
}
