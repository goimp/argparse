package argparse_test

import (
	"argparse"
	"testing"
)

func TestAction(t *testing.T) {
	action := &argparse.Action{
		OptionStrings: []string{"--verbose"},
		Dest:          "verbose",
		Nargs:         1,
		Default:       nil,
		Required:      false,
		Help:          "Enable verbose output",
		Metavar:       "VERBOSE",
	}

	// Test FormatUsage
	expectedUsage := "--verbose"
	if action.FormatUsage() != expectedUsage {
		t.Errorf("expected %q, got %q", expectedUsage, action.FormatUsage())
	}

	// Test GetKwargs
	kwargs := action.GetMap()
	if kwargs["Dest"] != "verbose" {
		t.Errorf("expected %q, got %q", "verbose", kwargs["Dest"])
	}
}
