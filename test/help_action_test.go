package argparse_test

import (
	"argparse"
	"fmt"
	"testing"
)

func TestHelpAction(t *testing.T) {

	p := &argparse.ArgumentParser{}

	ai := argparse.NewHelpAction(
		&argparse.Argument{
			OptionStrings: []string{"-h", "--help"},
			Dest:          "help",
			Nargs:         argparse.OPTIONAL,
			Required:      false,
			Help:          "Enable verbose output",
		},
	)

	fmt.Println("TestHelpAction:")
	prettyPrintMap(ai.GetMap())

	p.PrintHelp(nil)

}
