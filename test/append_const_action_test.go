package argparse_test

import (
	"argparse"
	"fmt"
	"reflect"
	"testing"
)

func TestAppendConstAction(t *testing.T) {

	n := argparse.NewNamespace(map[string]any{
		"foo": []any{"bar"},
	})

	ai := argparse.NewAppendConstAction(
		&argparse.Argument{
			OptionStrings: []string{"-f", "--foo"},
			Dest:          "foo",
			Const:         "baz",
			Default:       "baz",
			Help:          "Enable verbose output",
			Required:      false,
		},
	)

	fmt.Println("TestAction:")
	prettyPrintMap(ai.GetMap())

	ai.Call(nil, n, "baz", "")
	a := ai.Struct()

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

	if f := ai.FormatUsage(); !(f == "-f") {
		t.Errorf("Not found attribute %s in namespace\n", a.Dest)
	} else {
		fmt.Printf("Format usage: %s\n", f)
	}
}
