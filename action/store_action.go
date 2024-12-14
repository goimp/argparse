package action

import (
	"fmt"
	"reflect"
)

// StoreAction represents an action that stores a value in the namespace.
type StoreAction struct {
	OptionStrings []string
	Dest          string
	Nargs         int
	Const         interface{}
	Default       interface{}
	Type          interface{}
	Choices       []interface{}
	Required      bool
	Help          string
	Metavar       interface{}
	Deprecated    bool
}

// Constants for nargs values
const (
	OPTIONAL = -1
)

// NewStoreAction creates a new StoreAction with the provided parameters.
func NewStoreAction(optionStrings []string, dest string, nargs int, constVal interface{}, defaultVal interface{}, help string) (*StoreAction, error) {
	// Check if nargs is 0
	if nargs == 0 {
		return nil, fmt.Errorf("nargs for store actions must be != 0")
	}

	// Check if const is provided but nargs is not OPTIONAL
	if constVal != nil && nargs != OPTIONAL {
		return nil, fmt.Errorf("nargs must be %d to supply const", OPTIONAL)
	}

	return &StoreAction{
		OptionStrings: optionStrings,
		Dest:          dest,
		Nargs:         nargs,
		Const:         constVal,
		Default:       defaultVal,
		Help:          help,
	}, nil
}

// SetValue sets the value in the namespace (a struct or map) based on dest.
func (a *StoreAction) SetValue(namespace interface{}, value interface{}) error {
	// Ensure that the namespace is a pointer to a struct or map
	v := reflect.ValueOf(namespace)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("namespace must be a pointer")
	}

	// Handle setting value for struct or map
	if v.Elem().Kind() == reflect.Struct {
		field := v.Elem().FieldByName(a.Dest)
		if !field.IsValid() {
			return fmt.Errorf("no such field: %s in struct", a.Dest)
		}

		if !field.CanSet() {
			return fmt.Errorf("cannot set field: %s", a.Dest)
		}
		field.Set(reflect.ValueOf(value))
	} else if v.Elem().Kind() == reflect.Map {
		// For map types, we can set a key-value pair
		v.Elem().SetMapIndex(reflect.ValueOf(a.Dest), reflect.ValueOf(value))
	} else {
		return fmt.Errorf("namespace must be a struct or map")
	}

	return nil
}

// FormatUsage returns the formatted usage string.
func (a *StoreAction) FormatUsage() string {
	return fmt.Sprintf("Option: %v, Destination: %s", a.OptionStrings, a.Dest)
}

// SetValue sets the constant value in the namespace (a struct or map) based on dest.
func (a *StoreConstAction) SetValue(namespace interface{}) error {
	// Ensure that the namespace is a pointer to a struct or map
	v := reflect.ValueOf(namespace)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("namespace must be a pointer")
	}

	// Handle setting value for struct or map
	if v.Elem().Kind() == reflect.Struct {
		field := v.Elem().FieldByName(a.Dest)
		if !field.IsValid() {
			return fmt.Errorf("no such field: %s in struct", a.Dest)
		}

		if !field.CanSet() {
			return fmt.Errorf("cannot set field: %s", a.Dest)
		}
		field.Set(reflect.ValueOf(a.Const))
	} else if v.Elem().Kind() == reflect.Map {
		// For map types, we can set a key-value pair
		v.Elem().SetMapIndex(reflect.ValueOf(a.Dest), reflect.ValueOf(a.Const))
	} else {
		return fmt.Errorf("namespace must be a struct or map")
	}

	return nil
}

// FormatUsage returns the formatted usage string.
func (a *StoreConstAction) FormatUsage() string {
	return fmt.Sprintf("Option: %v, Constant: %v", a.OptionStrings, a.Const)
}
