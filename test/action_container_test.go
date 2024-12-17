package argparse_test

import (
	"argparse"
	"testing"
)

func TestActionsContainer(t *testing.T) {
	container := argparse.NewActionsContainer(
		"testDesc",
		"-",
		nil,
		nil,
	)

	container.AddArgument(
		&argparse.Argument{
			OptionStrings: []string{"-f", "--foo"},
			Action:        "store",
		},
	)

}
