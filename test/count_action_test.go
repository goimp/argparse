package argparse_test

import (
	"argparse"
	"fmt"
	"reflect"
	"testing"
)

func TestCountAction(t *testing.T) {

	n := argparse.NewNamespace(map[string]any{
		// "foo": 1,
	})

	ai := argparse.NewCountAction(
		&argparse.Argument{
			OptionStrings: []string{"-f", "--foo"},
			Dest:          "foo",
			Default:       "baz",
			Help:          "Enable verbose output",
			Required:      false,
		},
	)

	fmt.Println("TestCountAction:")
	prettyPrintMap(ai.GetMap())

	ai.Call(nil, n, 5, "")
	a := ai.Struct()

	referenceValue := 1

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
