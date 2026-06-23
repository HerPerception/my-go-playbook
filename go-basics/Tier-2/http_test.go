package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFetchURL_Success(t *testing.T) {
	t.Run("Successfully fetches content from a URL and reads body", func(t *testing.T) {
		expectedResponse := `{"message": "success", "status": "active"}`

		// Create a local mock HTTP server to simulate the network request safely
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(expectedResponse))
		}))
		// Shut down the mock server once the test concludes
		defer server.Close()

		// Execute your target function using the mock server's temporary local URL
		got, err := FetchURL(server.URL)
		if err != nil {
			t.Fatalf("FetchURL returned an unexpected error: %v", err)
		}

		if got != expectedResponse {
			t.Errorf("FetchURL() = %q; want %q", got, expectedResponse)
		}
	})
}

func TestFetchURL_NetworkError(t *testing.T) {
	t.Run("Propagates error safely when hitting an invalid URL", func(t *testing.T) {
		// An invalid URL format will cause http.Get to fail immediately
		invalidURL := "http://invalid-domain-that-does-not-exist-12345.xyz"

		// Execute target function (Must return an error and NOT panic)
		_, err := FetchURL(invalidURL)
		if err == nil {
			t.Fatal("Expected an error when attempting to fetch an invalid domain, but got nil")
		}
	})
}

func TestFetchURL_MainOutputTrimmingLogic(t *testing.T) {
	t.Run("Verification of long payload safety for byte slicing", func(t *testing.T) {
		// Create a mock server that sends a very large response body
		longPayload := strings.Repeat("A", 500) // 500 bytes of data
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(longPayload))
		}))
		defer server.Close()

		got, err := FetchURL(server.URL)
		if err != nil {
			t.Fatalf("FetchURL failed during length check scenario: %v", err)
		}

		// Ensure the function accurately returned the complete body string
		if len(got) != 500 {
			t.Errorf("Expected full response string length to be 500, but got %d", len(got))
		}

		// Simulating the requirement 5 slice operation from main(): print first 200 bytes
		trimmed := got[:200]
		if len(trimmed) != 200 {
			t.Errorf("Simulated main() slicing failed; got string length %d, want 200", len(trimmed))
		}
	})
}
