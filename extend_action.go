package argparse

import (
	"fmt"
	"reflect"
)

// ExtendAction represents an action that displays the version information.
type ExtendAction struct {
	*AppendAction
}

// NewExtendAction creates a new ExtendAction.
func NewExtendAction(argument *Argument) ActionInterface {
	action := NewAppendAction(argument)
	return &ExtendAction{AppendAction: action.(*AppendAction)}
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

	items = CopyItems(items)
	if !isSlice(items) {
		return fmt.Errorf("values is not a slice")
	}
	items = append(items.([]any), values.([]any)...)
	namespace.Set(a.Dest, items)
	return nil
}

// // Make sure StoreTrueAction implements ActionInterface
// func (a *ExtendAction) Struct() *Action {
// 	return a.AppendAction.Struct() // Call Struct() from StoreConstAction
// }

// func (a *ExtendAction) GetMap() map[string]any {
// 	return a.AppendAction.GetMap()
// }

// func (a *ExtendAction) FormatUsage() string {
// 	return a.AppendAction.FormatUsage()
// }

// func (a *ExtendAction) GetSubActions() []ActionInterface {
// 	return a.AppendAction.GetSubActions()
// }
