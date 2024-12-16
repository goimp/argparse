package action

import (
	"argparse"
	"argparse/copy_items"
	"argparse/namespace"
	"fmt"
)

// AppendAction represents an action that appends values to a slice.
type AppendAction struct {
	Action // Embedding Action to reuse functionality
}

// NewAppendAction creates a new AppendAction.
func NewAppendAction(
	optionStrings []string,
	dest string,
	nargs any,
	constVal any,
	defaultVal any,
	typ Type,
	choices []any,
	required bool,
	help string,
	metavar string,
	deprecated bool,
) (*AppendAction, error) {

	switch v := nargs.(type) {
	case int:
		if v == 0 {
			return nil, fmt.Errorf(
				"nargs for append actions must be != 0; if arg strings are not supplying the value to append, the append const action may be more appropriate",
			)
		}
	case string:
		if constVal != nil && v != argparse.OPTIONAL {
			return nil, fmt.Errorf("nargs must %s to supply const", argparse.OPTIONAL)
		}
	case nil:
		break
	default:
		return nil, fmt.Errorf("nargs must be an integer or a string literal (e.g., %s)", argparse.OPTIONAL)
	}

	return &AppendAction{
		Action: Action{
			OptionStrings: optionStrings,
			Dest:          dest,
			Nargs:         nargs,
			Const:         constVal,
			Default:       defaultVal,
			Type:          typ,
			Choices:       choices,
			Required:      required,
			Help:          help,
			Metavar:       metavar,
			Deprecated:    deprecated,
		},
	}, nil
}

func (a *AppendAction) Call(parser any, namespace *namespace.Namespace, values []any, optionString string) {
	items, found := namespace.Get(a.Dest)
	if !found {
		items = []any{}
	}
	items = copy_items.CopyItems(items)
	items = append(items.([]any), values...)
	namespace.Set(a.Dest, items)
}
