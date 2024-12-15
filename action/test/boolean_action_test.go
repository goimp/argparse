package action_test

import (
	"argparse/action"
	"argparse/namespace"
	"fmt"
	"testing"
)

func TestBooleanAction(t *testing.T) {

	n := namespace.NewNamespace(map[string]any{})

	a, err := action.NewBooleanOptionalAction(
		[]string{"-v", "--verbose"},
		"verbose",
		nil,
		false,
		"Enable verbose output",
		false,
	)

	if err != nil {
		t.Errorf("BooleanOptionalAction creation error: %s", err)
	}

	fmt.Printf("Kwargs: %v\n", a.GetKwargs())

	a.Call(nil, n, nil, "--verbose")

	if value, found := n.Get("verbose"); !found {
		t.Errorf("Not found attribute %s in namespace\n", a.Dest)
	} else {
		fmt.Printf("%s: %v\n", a.Dest, value)
	}

	if f := a.FormatUsage(); !(f == "-v | --verbose | --no-verbose") {
		t.Errorf("Not found attribute %s in namespace\n", a.Dest)
	} else {
		fmt.Printf("Format usage: %s\n", f)
	}
}
