package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestLoggerInterfaceAndImplementations(t *testing.T) {
	// 1. Compile-Time Interface Verification Check
	t.Run("Interface Assertions", func(t *testing.T) {
		var _ Logger = ConsoleLogger{}
		var _ Logger = &FileLogger{} // Checked as pointer since it manages file state/operations
	})

	// 2. Test ConsoleLogger by capturing the system standard output (stdout)
	t.Run("ConsoleLogger Outputs To Stdout", func(t *testing.T) {
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		logger := ConsoleLogger{}
		testMessage := "testing console logger output"
		logger.Log(testMessage)

		w.Close()
		os.Stdout = oldStdout

		var buf bytes.Buffer
		_, _ = io.Copy(&buf, r)
		output := buf.String()

		if !strings.Contains(output, testMessage) {
			t.Errorf("ConsoleLogger expected to log %q, but stdout got %q", testMessage, output)
		}
	})

	// 3. Test FileLogger by checking if it accurately creates and writes to a disk file
	t.Run("FileLogger Writes Content To Disk File", func(t *testing.T) {
		tempFile := "test_run_output.log"

		// Ensure clean state before running test
		_ = os.Remove(tempFile)
		defer os.Remove(tempFile)

		logger := &FileLogger{Filename: tempFile}
		testMessage := "testing file logger output"
		logger.Log(testMessage)

		// Read back what was written to verify content
		content, err := os.ReadFile(tempFile)
		if err != nil {
			t.Fatalf("Failed to read expected log file: %v", err)
		}

		if !strings.Contains(string(content), testMessage) {
			t.Errorf("FileLogger expected to find text %q, but file contained %q", testMessage, string(content))
		}
	})

	// 4. Test the RunApp framework function using an isolated mock strategy
	t.Run("RunApp Invokes Log Method Properly", func(t *testing.T) {
		mock := &mockLogger{}

		RunApp(mock)

		expectedMessage := "app started"
		if mock.lastMessage != expectedMessage {
			t.Errorf("RunApp called Log with %q; want %q", mock.lastMessage, expectedMessage)
		}
	})
}

// Custom internal mock helper to isolate testing the RunApp logic without hitting disk or stdout
type mockLogger struct {
	lastMessage string
}

func (m *mockLogger) Log(message string) {
	m.lastMessage = message
}
