package argparse_test

import (
	"testing"
)

func TestCountAction(t *testing.T) {

	// n := argparse.NewNamespace(map[string]any{
	// 	// "foo": 1,
	// })

	// a, err := argparse.NewCountAction(
	// 	[]string{"-f", "--foo"},
	// 	"foo",
	// 	0,
	// 	false,
	// 	"Enable verbose output",
	// 	false,
	// )

	// if err != nil {
	// 	t.Errorf("CountAction creation error: %s", err)
	// }

	// fmt.Printf("Kwargs: %v\n", a.GetMap())

	// a.Call(nil, n, 5, "")

	// referenceValue := 1

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
