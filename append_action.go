package argparse

import (
	"fmt"
)

// AppendAction represents an action that appends values to a slice.
type AppendAction struct {
	*Action // Embedding Action to reuse functionality
}

// NewAppendAction creates a new AppendAction.
func NewAppendAction(argument *Argument) ActionInterface {

	switch v := argument.Nargs.(type) {
	case int:
		if v == 0 {
			panic(
				"nargs for append actions must be != 0; if arg strings are not supplying the value to append, the append const action may be more appropriate",
			)
		}
	case string:
		if argument.Const != nil && v != OPTIONAL {
			panic(fmt.Sprintf("nargs must %s to supply const", OPTIONAL))
		}
	case nil:
		break
	default:
		panic(fmt.Sprintf("nargs must be an integer or a string literal (e.g., %s)", OPTIONAL))
	}

	return &AppendAction{
		Action: &Action{
			OptionStrings: argument.OptionStrings,
			Dest:          argument.Dest,
			Nargs:         argument.Nargs,
			Const:         argument.Const,
			Default:       argument.Default,
			Type:          argument.Type,
			Choices:       argument.Choices,
			Required:      argument.Required,
			Help:          argument.Help,
			MetaVar:       argument.MetaVar,
			Deprecated:    argument.Deprecated,
		},
	}
}

func (a *AppendAction) Call(parser *ArgumentParser, namespace *Namespace, values any, optionString string) error {
	items, found := namespace.Get(a.Dest)
	if !found {
		items = []any{}
	}
	items = CopyItems(items)
	items = append(items.([]any), values)
	namespace.Set(a.Dest, items)
	return nil
}
