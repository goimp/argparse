package argparse_test

import (
	"argparse"
	"argparse/argument_parser"
	"fmt"
	"testing"
)

func TestHelpAction(t *testing.T) {

	p := &argument_parser.ArgumentParser{}

	a, err := argparse.NewHelpAction(
		[]string{"-h", "--help"},
		"help",
		nil,
		"Enable verbose output",
		false,
	)

	if err != nil {
		t.Errorf("HelpAction creation error: %s", err)
	}

	fmt.Printf("Kwargs: %v\n", a.GetKwargs())

	p.PrintHelp(nil)

}
