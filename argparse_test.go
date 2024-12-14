package argparse

import (
	"testing"
)

// TestGreet tests the Greet function.
func TestGreet(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "Greet John", input: "John", expected: "Hello, John!"},
		{name: "Greet empty string", input: "", expected: "Hello, !"},
		{name: "Greet special characters", input: "@User", expected: "Hello, @User!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Greet(tt.input)
			if got != tt.expected {
				t.Errorf("Greet(%q) = %q; want %q", tt.input, got, tt.expected)
			}
		})
	}
}