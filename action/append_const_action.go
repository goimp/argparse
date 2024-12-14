package action

import (
	"fmt"
	"reflect"
)

// AppendConstAction represents an action that appends a constant value to a slice.
type AppendConstAction struct {
	Action        // Embedding Action to reuse functionality
	Const         interface{}
	Dest          string
	OptionStrings []string
}

// NewAppendConstAction creates a new AppendConstAction.
func NewAppendConstAction(optionStrings []string, dest string, constVal interface{}, help string) *AppendConstAction {
	return &AppendConstAction{
		Action:        Action{OptionStrings: optionStrings, Dest: dest},
		Const:         constVal,
		Dest:          dest,
		OptionStrings: optionStrings,
	}
}

// SetValue appends the constant value to the specified destination in the namespace.
func (a *AppendConstAction) SetValue(namespace interface{}) error {
	destField := reflect.ValueOf(namespace).Elem().FieldByName(a.Dest)

	if destField.IsValid() && destField.CanSet() {
		// Check if it's a slice and append the constant value
		if destField.Kind() == reflect.Slice {
			destField.Set(reflect.Append(destField, reflect.ValueOf(a.Const)))
		} else {
			// Handle error if the field is not a slice
			return fmt.Errorf("destination %s is not a slice", a.Dest)
		}
	} else {
		return fmt.Errorf("destination %s not found or cannot be set", a.Dest)
	}

	return nil
}
