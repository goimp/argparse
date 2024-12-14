package action

import (
	"fmt"
	"reflect"
)

// AppendAction represents an action that appends values to a slice.
type AppendAction struct {
	Action        // Embedding Action to reuse functionality
	Default       interface{}
	Dest          string
	OptionStrings []string
}

// NewAppendAction creates a new AppendAction.
func NewAppendAction(optionStrings []string, dest string, defaultVal interface{}, help string) *AppendAction {
	return &AppendAction{
		Action:        Action{OptionStrings: optionStrings, Dest: dest},
		Default:       defaultVal,
		Dest:          dest,
		OptionStrings: optionStrings,
	}
}

// SetValue appends the value to the specified destination in the namespace.
func (a *AppendAction) SetValue(namespace interface{}, value interface{}) error {
	destField := reflect.ValueOf(namespace).Elem().FieldByName(a.Dest)

	if destField.IsValid() && destField.CanSet() {
		// Check if it's a slice and append the value
		if destField.Kind() == reflect.Slice {
			destField.Set(reflect.Append(destField, reflect.ValueOf(value)))
		} else {
			// Handle error if the field is not a slice
			return fmt.Errorf("destination %s is not a slice", a.Dest)
		}
	} else {
		return fmt.Errorf("destination %s not found or cannot be set", a.Dest)
	}

	return nil
}
