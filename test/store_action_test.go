package argparse_test

import (
	"argparse"
	"argparse/namespace"
	"fmt"
	"testing"
)

func TestStoreAction(t *testing.T) {

	n := namespace.NewNamespace(map[string]any{})

	a, err := argparse.NewStoreAction(
		[]string{"-f", "--foo"},
		"foo",
		argparse.OPTIONAL,
		nil,
		nil,
		nil,
		nil,
		false,
		"Enable verbose output",
		"",
		false,
	)

	if err != nil {
		t.Errorf("StoreAction creation error: %s", err)
	}

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
