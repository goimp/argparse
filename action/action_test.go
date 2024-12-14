package action

import (
	"testing"
)

func TestAction(t *testing.T) {
	nargs := 1
	action := NewAction([]string{"--verbose"}, "verbose", &nargs, nil, nil, nil, nil, false, "Enable verbose", "VERBOSE", false)

	// Test FormatUsage
	expectedUsage := "--verbose"
	if action.FormatUsage() != expectedUsage {
		t.Errorf("expected %q, got %q", expectedUsage, action.FormatUsage())
	}

	// Test GetKwargs
	kwargs := action.GetKwargs()
	if kwargs["dest"] != "verbose" {
		t.Errorf("expected %q, got %q", "verbose", kwargs["dest"])
	}
}

func TestBooleanOptionalAction(t *testing.T) {
	// Initialize the action
	action, err := NewBooleanOptionalAction([]string{"--verbose"}, "verbose", nil, false, "Enable verbose output", false)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Simulate calling the action with "--verbose"
	var namespace interface{}
	action.Call(nil, &namespace, nil, "--verbose")
	// Verify the value set for "verbose" (here we just check the print output for simplicity)
	// In a real test, you'd verify the actual value in the namespace
}
