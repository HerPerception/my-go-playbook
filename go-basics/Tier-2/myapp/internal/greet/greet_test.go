package greet

import "testing"

func TestHello(t *testing.T) {
	// Table-driven test setup to cover basic scenarios
	tests := []struct {
		name       string // The description of the subtest
		inputName  string // The argument sent to Hello()
		wantOutput string // The exact expected return string
	}{
		{
			name:       "Standard name greeting",
			inputName:  "Alice",
			wantOutput: "Hello, Alice!",
		},
		{
			name:       "Empty string handling",
			inputName:  "",
			wantOutput: "Hello, !",
		},
		{
			name:       "Name with whitespace",
			inputName:  "Bob Smith",
			wantOutput: "Hello, Bob Smith!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOutput := Hello(tt.inputName)

			if gotOutput != tt.wantOutput {
				t.Errorf("Hello(%q) = %q; want %q", tt.inputName, gotOutput, tt.wantOutput)
			}
		})
	}
}
