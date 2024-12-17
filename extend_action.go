package argparse

import (
	"argparse/copy_items"
	"fmt"
	"reflect"
)

// ExtendAction represents an action that displays the version information.
type ExtendAction struct {
	*AppendAction
}

// NewExtendAction creates a new ExtendAction.
func NewExtendAction(argument *Argument) *ExtendAction {
	action := NewAppendAction(argument)
	return &ExtendAction{AppendAction: action}
}

func isSlice(value any) bool {
	return reflect.TypeOf(value).Kind() == reflect.Slice
}

// SetValue prints the version information and exits the program.
func (a *ExtendAction) Call(parser *ArgumentParser, namespace *Namespace, values any, optionString string) error {
	items, found := namespace.Get(a.Dest)
	if !found {
		items = []any{}
	}

	items = copy_items.CopyItems(items)
	if !isSlice(items) {
		return fmt.Errorf("values is not a slice")
	}
	items = append(items.([]any), values.([]any)...)
	namespace.Set(a.Dest, items)
	return nil
}
