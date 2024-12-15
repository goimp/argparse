package action_test

import (
	"argparse/action"
	"argparse/namespace"
	"fmt"
	"testing"
)

func TestStoreFalseAction(t *testing.T) {

	n := namespace.NewNamespace(map[string]any{})

	a, err := action.NewStoreFalseAction(
		[]string{"-f", "--foo"},
		"foo",
		true,
		false,
		"Enable verbose output",
		false,
	)

	if err != nil {
		t.Errorf("StoreFalseAction creation error: %s", err)
	}

	fmt.Printf("Kwargs: %v\n", a.GetKwargs())

	a.Call(nil, n, []any{1, 2, 3}, "")

	if value, found := n.Get("foo"); !found {
		t.Errorf("Not found attribute %s in namespace\n", a.Dest)
	} else {
		if value != a.Const {
			t.Errorf("Action value error, expected %v, got %v\n", a.Const, value)
		} else {
			fmt.Printf("%s: %v\n", a.Dest, value)
		}
	}

	if f := a.FormatUsage(); !(f == "-f") {
		t.Errorf("Not found attribute %s in namespace\n", a.Dest)
	} else {
		fmt.Printf("Format usage: %s\n", f)
	}
}
