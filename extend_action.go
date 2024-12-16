package argparse

import (
	"argparse/copy_items"
	"argparse/namespace"
)

// ExtendAction represents an action that displays the version information.
type ExtendAction struct {
	AppendAction
}

// NewExtendAction creates a new ExtendAction.
func NewExtendAction(
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
) (*ExtendAction, error) {
	action, error := NewAppendAction(
		optionStrings,
		dest,
		nargs,
		constVal,
		defaultVal,
		typ,
		choices,
		required,
		help,
		metavar,
		deprecated,
	)
	if error != nil {
		return nil, error
	}
	return &ExtendAction{
		AppendAction: *action,
	}, nil
}

// SetValue prints the version information and exits the program.
func (a *ExtendAction) Call(parser *ArgumentParser, namespace *namespace.Namespace, values []any, optionString string) {
	items, found := namespace.Get(a.Dest)
	if !found {
		items = []any{}
	}
	items = copy_items.CopyItems(items)
	items = append(items.([]any), values...)
	namespace.Set(a.Dest, items)
}
