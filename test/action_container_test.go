package argparse_test

import (
	"argparse"
	"fmt"
	"reflect"
	"testing"
)

func prettyPrintMap(mapping map[string]any) {
	for key, value := range mapping {
		switch t := value.(type) {
		case int:
			fmt.Printf("%-10s\t:\t%d\n", key, t)
		case float32, float64:
			fmt.Printf("%-10s\t:\t%.6f\n", key, t)
		case string:
			fmt.Printf("%-10s\t:\t%s\n", key, t)
		default:
			fmt.Printf("%-10s\t:\t%v\n", key, t)
		}
	}
	fmt.Println()
}

func TestActionsContainer(t *testing.T) {
	container := argparse.NewActionsContainer(
		"testDesc",
		"-",
		nil,
		nil,
	)

	container.AddArgument(
		&argparse.Argument{
			OptionStrings: []string{"-f", "--foo"},
			Action:        "store",
		},
	)

	container.AddArgument(
		&argparse.Argument{
			OptionStrings: []string{"-b", "--bar"},
			Action:        "store_true",
		},
	)

	container.AddArgument(
		&argparse.Argument{
			OptionStrings: []string{"-z"},
			Action:        "store_false",
		},
	)

	container.AddArgument(
		&argparse.Argument{
			OptionStrings: []string{"-c", "-const"},
			Action:        "store_const",
			Dest:          "constant",
		},
	)

	container.AddArgument(
		&argparse.Argument{
			OptionStrings: []string{"--counter"},
			Action:        "count",
			Dest:          "counter",
		},
	)

	for _, actIntf := range container.Actions {
		fmt.Println(reflect.TypeOf(actIntf))
		prettyPrintMap(actIntf.GetKwargs())
	}
}
