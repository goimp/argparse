package argparse_test

import (
	"fmt"
	"testing"

	"github.com/goimp/argparse"
)

func TestBooleanAction(t *testing.T) {

	n := argparse.NewNamespace(map[string]any{})

	ai := argparse.NewBooleanOptionalAction(
		&argparse.Argument{
			OptionStrings: []string{"-v", "--verbose"},
			Dest:          "verbose",
			Help:          "Enable verbose output",
			Required:      false,
		},
	)

	fmt.Println("TestBooleanAction:")
	prettyPrintMap(ai.GetMap())

	ai.Call(nil, n, nil, "--verbose")
	a := ai.Struct()

	if value, found := n.Get("verbose"); !found {
		t.Errorf("Not found attribute %s in namespace\n", a.Dest)
	} else {
		fmt.Printf("%s: %v\n", a.Dest, value)
	}

	if f := ai.FormatUsage(); !(f == "-v | --verbose | --no-verbose") {
		t.Errorf("Not found attribute %s in namespace\n", a.Dest)
	} else {
		fmt.Printf("Format usage: %s\n", f)
	}
}
