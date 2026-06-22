package main

import (
	"errors"
	"strings"
	"testing"
)

func TestValidationError_InterfaceImplementation(t *testing.T) {
	t.Run("Pointer strictly satisfies the standard error interface", func(t *testing.T) {
		// Validates that *ValidationError satisfies Go's built-in error contract
		var _ error = &ValidationError{}
	})
}

func TestValidationError_ErrorMessageFormat(t *testing.T) {
	t.Run("Error string contains both field and clear message details", func(t *testing.T) {
		err := &ValidationError{
			Field:   "Email",
			Message: "cannot be empty",
		}

		gotStr := err.Error()

		// Verify the custom error string output contains the critical context
		if !strings.Contains(gotStr, "Email") || !strings.Contains(gotStr, "cannot be empty") {
			t.Errorf("Error() output %q does not properly communicate field or message content", gotStr)
		}
	})
}

func TestValidateAge(t *testing.T) {
	tests := []struct {
		name        string
		age         int
		expectNil   bool
		wantField   string
		wantMessage string
	}{
		{
			name:      "Valid age minimum boundary",
			age:       0,
			expectNil: true,
		},
		{
			name:      "Valid standard adult age",
			age:       25,
			expectNil: true,
		},
		{
			name:      "Valid age maximum boundary",
			age:       150,
			expectNil: true,
		},
		{
			name:        "Invalid negative age triggers error",
			age:         -1,
			expectNil:   false,
			wantField:   "age",
			wantMessage: "out of range", // or whatever custom phrase you chose
		},
		{
			name:        "Invalid extreme high age triggers error",
			age:         151,
			expectNil:   false,
			wantField:   "age",
			wantMessage: "out of range",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateAge(tt.age)

			if tt.expectNil {
				if err != nil {
					t.Fatalf("ValidateAge(%d) expected no error, but got: %v", tt.age, err)
				}
				return // Test case complete for valid values
			}

			if err == nil {
				t.Fatalf("ValidateAge(%d) expected an error, but got nil", tt.age)
			}

			// Validate requirement 4: errors.As unwrapping contract
			var ve *ValidationError
			if !errors.As(err, &ve) {
				t.Fatalf("ValidateAge(%d) returned error that failed to unwrap via errors.As into a *ValidationError", tt.age)
			}

			// Confirm underlying field fields match exact target expectations
			if ve.Field != tt.wantField {
				t.Errorf("errors.As extracted Field = %q; want %q", ve.Field, tt.wantField)
			}

			// Simple verification that message string has content if match text varies
			if ve.Message == "" {
				t.Error("errors.As extracted ValidationError contains an empty Message field")
			}
		})
	}
}
