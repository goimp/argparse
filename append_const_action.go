package argparse

import (
	"argparse/copy_items"
)

// AppendConstAction represents an action that appends a constant value to a slice.
type AppendConstAction struct {
	*Action // Embedding Action to reuse functionality
}

// NewAppendConstAction creates a new AppendConstAction.
func NewAppendConstAction(argument *Argument) *AppendConstAction {
	return &AppendConstAction{
		Action: &Action{
			OptionStrings: argument.OptionStrings,
			Dest:          argument.Dest,
			Nargs:         0,
			Const:         argument.Const,
			Default:       argument.Default,
			Required:      argument.Required,
			Help:          argument.Help,
			Metavar:       argument.Metavar,
			Deprecated:    argument.Deprecated,
		},
	}
}

func (a *AppendConstAction) Call(parser *ArgumentParser, namespace *Namespace, values any, optionString string) error {
	items, found := namespace.Get(a.Dest)
	if !found {
		items = []any{}
	}
	items = copy_items.CopyItems(items)
	items = append(items.([]any), a.Const)
	namespace.Set(a.Dest, items)
	return nil
}
