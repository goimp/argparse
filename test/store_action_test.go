package argparse_test

import (
	"argparse"
	"fmt"
	"testing"
)

func TestStoreAction(t *testing.T) {

	n := argparse.NewNamespace(map[string]any{})

	ai := argparse.NewStoreAction(
		&argparse.Argument{
			OptionStrings: []string{"-f", "--foo"},
			Dest:          "foo",
			Nargs:         argparse.OPTIONAL,
			Required:      false,
			Help:          "Enable verbose output",
		},
	)

	fmt.Println("TestStoreConstAction:")
	prettyPrintMap(ai.GetMap())

	ai.Call(nil, n, []any{1, 2, 3}, "")

	action := ai.Struct()

	if value, found := n.Get("foo"); !found {
		t.Errorf("Not found attribute %s in namespace\n", action.Dest)
	} else {
		fmt.Printf("%s: %v\n", action.Dest, value)
	}

	if f := ai.FormatUsage(); !(f == "-f") {
		t.Errorf("Not found attribute %s in namespace\n", action.Dest)
	} else {
		fmt.Printf("Format usage: %s\n", f)
	}
}
