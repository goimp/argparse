package argparse_test

import (
	"argparse"
	"testing"
)

func TestArgumentError(t *testing.T) {
	// Test with valid argument
	arg := &argparse.Action{
		OptionStrings: []string{"--test"},
	}
	err := argparse.NewArgumentError(arg, "invalid option")
	expected := "argument --test: invalid option"
	if err.Error() != expected {
		t.Errorf("expected %q, got %q", expected, err.Error())
	}

	// Test with no argument
	err = argparse.NewArgumentError(nil, "missing argument")
	expected = "missing argument"
	if err.Error() != expected {
		t.Errorf("expected %q, got %q", expected, err.Error())
	}

	// Test with metavar
	arg = &argparse.Action{
		MetaVar: "FILE",
	}
	err = argparse.NewArgumentError(arg, "file not found")
	expected = "argument FILE: file not found"
	if err.Error() != expected {
		t.Errorf("expected %q, got %q", expected, err.Error())
	}
}

func TestArgumentTypeError(t *testing.T) {
	err := argparse.NewArgumentTypeError("failed to convert type")
	expected := "failed to convert type"
	if err.Error() != expected {
		t.Errorf("expected %q, got %q", expected, err.Error())
	}
}
