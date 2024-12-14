package argument_error

import (
	"argparse/action"
	"testing"
)

func TestArgumentError(t *testing.T) {
	// Test with valid argument
	arg := &action.Argument{
		OptionStrings: []string{"--test"},
	}
	err := NewArgumentError(arg, "invalid option")
	expected := "argument --test: invalid option"
	if err.Error() != expected {
		t.Errorf("expected %q, got %q", expected, err.Error())
	}

	// Test with no argument
	err = NewArgumentError(nil, "missing argument")
	expected = "missing argument"
	if err.Error() != expected {
		t.Errorf("expected %q, got %q", expected, err.Error())
	}

	// Test with metavar
	arg = &action.Argument{
		Metavar: "FILE",
	}
	err = NewArgumentError(arg, "file not found")
	expected = "argument FILE: file not found"
	if err.Error() != expected {
		t.Errorf("expected %q, got %q", expected, err.Error())
	}
}

func TestArgumentTypeError(t *testing.T) {
	err := NewArgumentTypeError("failed to convert type")
	expected := "failed to convert type"
	if err.Error() != expected {
		t.Errorf("expected %q, got %q", expected, err.Error())
	}
}
