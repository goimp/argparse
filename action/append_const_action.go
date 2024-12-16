package action

import (
	"argparse/copy_items"
	"argparse/namespace"
)

// AppendConstAction represents an action that appends a constant value to a slice.
type AppendConstAction struct {
	Action // Embedding Action to reuse functionality
}

// NewAppendConstAction creates a new AppendConstAction.
func NewAppendConstAction(
	optionStrings []string,
	dest string,
	constVal any,
	defaultVal any,
	required bool,
	help string,
	metavar string,
	deprecated bool,
) (*AppendConstAction, error) {
	return &AppendConstAction{
		Action: Action{
			OptionStrings: optionStrings,
			Dest:          dest,
			Nargs:         0,
			Const:         constVal,
			Default:       defaultVal,
			Required:      required,
			Help:          help,
			Metavar:       metavar,
			Deprecated:    deprecated,
		},
	}, nil
}

func (a *AppendConstAction) Call(parser any, namespace *namespace.Namespace, values any, optionString string) {
	items, found := namespace.Get(a.Dest)
	if !found {
		items = []any{}
	}
	items = copy_items.CopyItems(items)
	items = append(items.([]any), a.Const)
	namespace.Set(a.Dest, items)
}
