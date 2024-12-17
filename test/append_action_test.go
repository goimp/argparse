package argparse_test

import (
	"testing"
)

func TestAppendAction(t *testing.T) {

	// n := argparse.NewNamespace(map[string]any{
	// 	"foo": []any{1, 2},
	// })

	// a, err := argparse.NewAppendAction(
	// 	[]string{"-f", "--foo"},
	// 	"foo",
	// 	argparse.OPTIONAL,
	// 	nil,
	// 	nil,
	// 	nil,
	// 	nil,
	// 	false,
	// 	"Enable verbose output",
	// 	"",
	// 	false,
	// )

	// if err != nil {
	// 	t.Errorf("AppendAction creation error: %s", err)
	// }

	// fmt.Printf("Kwargs: %v\n", a.GetMap())

	// a.Call(nil, n, []any{1, 2, 3}, "")

	// referenceValue := []any{1, 2, []any{1, 2, 3}}

	// if value, found := n.Get("foo"); !found {
	// 	t.Errorf("Not found attribute %s in namespace\n", a.Dest)
	// } else {
	// 	// Use reflect.DeepEqual to compare the values
	// 	if !reflect.DeepEqual(value, referenceValue) {
	// 		t.Errorf("Wrong value got, expected %v, got %v\n", referenceValue, value)
	// 	} else {
	// 		fmt.Printf("%s: %v\n", a.Dest, value)
	// 	}
	// }

	// if f := a.FormatUsage(); !(f == "-f") {
	// 	t.Errorf("Not found attribute %s in namespace\n", a.Dest)
	// } else {
	// 	fmt.Printf("Format usage: %s\n", f)
	// }
}
