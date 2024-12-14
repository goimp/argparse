package action

import (
	"fmt"
	"reflect"
)

// CountAction represents an action that counts the occurrences of a command-line option.
type CountAction struct {
	Action        // Embedding Action to reuse functionality
	Dest          string
	OptionStrings []string
}

// NewCountAction creates a new CountAction.
func NewCountAction(optionStrings []string, dest string, help string) *CountAction {
	return &CountAction{
		Action:        Action{OptionStrings: optionStrings, Dest: dest},
		Dest:          dest,
		OptionStrings: optionStrings,
	}
}

// SetValue increments the counter value in the namespace.
func (a *CountAction) SetValue(namespace interface{}) error {
	destField := reflect.ValueOf(namespace).Elem().FieldByName(a.Dest)

	if destField.IsValid() && destField.CanSet() {
		// Check if it's an integer field
		if destField.Kind() == reflect.Int {
			// Increment the counter
			destField.SetInt(destField.Int() + 1)
		} else {
			// Handle error if the field is not an integer
			return fmt.Errorf("destination %s is not an integer", a.Dest)
		}
	} else {
		return fmt.Errorf("destination %s not found or cannot be set", a.Dest)
	}

	return nil
}
