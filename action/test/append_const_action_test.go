package action_test

import (
	"argparse/action"
	"argparse/namespace"
	"fmt"
	"reflect"
	"testing"
)

func TestAppendConstAction(t *testing.T) {

	n := namespace.NewNamespace(map[string]any{
		"foo": []any{"bar"},
	})

	a, err := action.NewAppendConstAction(
		[]string{"-f", "--foo"},
		"foo",
		"baz",
		nil,
		false,
		"Enable verbose output",
		"",
		false,
	)

	if err != nil {
		t.Errorf("AppendConstAction creation error: %s", err)
	}

	fmt.Printf("Kwargs: %v\n", a.GetKwargs())

	a.Call(nil, n, "baz", "")

	referenceValue := []any{"bar", "baz"}

	if value, found := n.Get("foo"); !found {
		t.Errorf("Not found attribute %s in namespace\n", a.Dest)
	} else {
		// Use reflect.DeepEqual to compare the values
		if !reflect.DeepEqual(value, referenceValue) {
			t.Errorf("Wrong value got, expected %v, got %v\n", referenceValue, value)
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
