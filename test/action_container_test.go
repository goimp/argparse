package argparse_test

import (
	"fmt"
	"reflect"
	"sort"
	"testing"

	"github.com/goimp/argparse"
)

func prettyPrintMap(mapping map[string]any) {
	// Крок 1: Зібрати ключі у зріз
	keys := make([]string, 0, len(mapping))
	for key := range mapping {
		keys = append(keys, key)
	}

	// Крок 2: Відсортувати ключі
	sort.Strings(keys)

	// Крок 3: Ітерувати по відсортованих ключах і виводити значення
	for _, key := range keys {
		value := mapping[key]
		switch t := value.(type) {
		case int:
			fmt.Printf("%-10s\t:\t%d\n", key, t)
		case float32:
			fmt.Printf("%-10s\t:\t%.6f\n", key, t)
		case float64:
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
			Dest:          "store",
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

	container.AddArgument(
		&argparse.Argument{
			OptionStrings: []string{"--help"},
			Action:        "help",
			Dest:          "help",
		},
	)

	container.AddArgument(
		&argparse.Argument{
			OptionStrings: []string{"-V", "--version"},
			Action:        "version",
			Help:          "Print version",
		},
	)

	container.AddArgument(
		&argparse.Argument{
			OptionStrings: []string{"--append"},
			Action:        "append",
		},
	)

	container.AddArgument(
		&argparse.Argument{
			OptionStrings: []string{"--extend"},
			Action:        "extend",
		},
	)

	container.AddArgument(
		&argparse.Argument{
			OptionStrings: []string{"--append-const"},
			Action:        "append_const",
		},
	)

	// Positional
	container.AddArgument(
		&argparse.Argument{
			OptionStrings: []string{"pos"},
			Action:        "store",
			Nargs:         2,
		},
	)

	container.AddArgument(
		&argparse.Argument{
			OptionStrings: []string{},
			Dest:          "pos2",
			Action:        "store",
			Nargs:         1,
		},
	)

	container.Struct().SetDefaults(
		map[string]any{
			"default": "<Default value for store>",
		},
	)

	defAct := container.AddArgument(
		&argparse.Argument{
			OptionStrings: []string{"--def"},
			Dest:          "default",
			Action:        "store",
			Nargs:         1,
		},
	)

	for _, actIntf := range container.(*argparse.ActionsContainer).Actions {
		fmt.Println(reflect.TypeOf(actIntf))
		prettyPrintMap(actIntf.GetMap())
	}

	if container.Struct().GetDefault("default").(string) != "<Default value for store>" {
		t.Errorf("Wrong default value")
	}

	container.Struct().RemoveAction(defAct)

}
