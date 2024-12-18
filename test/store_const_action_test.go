package argparse_test

import (
	"fmt"
	"testing"

	"github.com/goimp/argparse"
)

func TestStoreConstAction(t *testing.T) {

	n := argparse.NewNamespace(map[string]any{})

	ai := argparse.NewStoreConstAction(
		&argparse.Argument{
			OptionStrings: []string{"-f", "--foo"},
			Dest:          "foo",
			Default:       "bar",
		},
	)

	fmt.Println("TestStoreConstAction:")
	prettyPrintMap(ai.GetMap())

	ai.Call(nil, n, []any{1, 2, 3}, "")
	a := ai.Struct()

	if value, found := n.Get("foo"); !found {
		t.Errorf("Not found attribute %s in namespace\n", a.Dest)
	} else {
		if value != a.Const {
			t.Errorf("Action value error, expected %v, got %v\n", a.Const, value)
		} else {
			fmt.Printf("%s: %v\n", a.Dest, value)
		}
	}

	if f := ai.FormatUsage(); !(f == "-f") {
		t.Errorf("Not found attribute %s in namespace\n", a.Dest)
	} else {
		fmt.Printf("Format usage: %s\n", f)
	}
}
