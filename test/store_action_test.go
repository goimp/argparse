package argparse_test

import (
	"argparse"
	"fmt"
	"testing"
)

func TestStoreAction(t *testing.T) {

	n := argparse.NewNamespace(map[string]any{})

	a := argparse.NewStoreAction(
		&argparse.Argument{
			OptionStrings: []string{"-f", "--foo"},
			Dest:          "foo",
			Nargs:         argparse.OPTIONAL,
			Required:      false,
			Help:          "Enable verbose output",
		},
	)

	fmt.Printf("Kwargs: %v\n", a.GetKwargs())

	a.Call(nil, n, []any{1, 2, 3}, "")

	if value, found := n.Get("foo"); !found {
		t.Errorf("Not found attribute %s in namespace\n", a.Dest)
	} else {
		fmt.Printf("%s: %v\n", a.Dest, value)
	}

	if f := a.FormatUsage(); !(f == "-f") {
		t.Errorf("Not found attribute %s in namespace\n", a.Dest)
	} else {
		fmt.Printf("Format usage: %s\n", f)
	}
}
