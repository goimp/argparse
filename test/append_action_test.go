package argparse_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/goimp/argparse"
)

func TestAppendAction(t *testing.T) {

	n := argparse.NewNamespace(map[string]any{
		"foo": []any{1, 2},
	})

	ai := argparse.NewAppendAction(
		&argparse.Argument{
			OptionStrings: []string{"-f", "--foo"},
			Dest:          "foo",
			Default:       "baz",
			Help:          "Enable verbose output",
			Required:      false,
		},
		// []string{"-f", "--foo"},
		// "foo",
		// argparse.OPTIONAL,
		// nil,
		// nil,
		// nil,
		// nil,
		// false,
		// "Enable verbose output",
		// "",
		// false,
	)

	fmt.Println("TestBooleanAction:")
	prettyPrintMap(ai.GetMap())

	ai.Call(nil, n, []any{1, 2, 3}, "")
	a := ai.Struct()

	referenceValue := []any{1, 2, []any{1, 2, 3}}

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
