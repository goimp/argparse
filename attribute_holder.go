package argparse

import (
	"fmt"
	"reflect"
	"strings"
)

// AttributeHolder provides a representation of a struct in the format:
// StructName(attr=value, attr=value, ...)

type AttributeHolderInterface_ interface {
	Repr() string
	GetArgs() []string
	GetKwargs(v reflect.Value)
}

type AttributeHolder_ struct{}

// Repr returns a string representation of the struct in the format:
// StructName(attr=value, attr=value, ...)
func (ah *AttributeHolder_) Repr() string {
	// Get the type and value of the current struct
	t := reflect.TypeOf(ah).Elem()
	v := reflect.ValueOf(ah).Elem()

	// Get the struct name
	structName := t.Name()

	// Collect arguments
	var argStrings []string
	starArgs := make(map[string]any)

	// Handle args and kwargs
	argStrings = append(argStrings, ah.GetArgs()...)
	for key, value := range ah.GetKwargs(v) {
		if isValidIdentifier(key) {
			argStrings = append(argStrings, fmt.Sprintf("%s=%v", key, value))
		} else {
			starArgs[key] = value
		}
	}

	// Add starArgs if any
	if len(starArgs) > 0 {
		argStrings = append(argStrings, fmt.Sprintf("**%v", starArgs))
	}

	return fmt.Sprintf("%s(%s)", structName, strings.Join(argStrings, ", "))
}

// GetArgs returns the positional arguments of the struct.
// By default, it returns an empty list. Override in derived structs if needed.
func (ah *AttributeHolder_) GetArgs() []string {
	return []string{}
}

// GetKwargs returns the keyword arguments of the struct.
// It collects all exported fields of the struct as key-value pairs.
func (ah *AttributeHolder_) GetKwargs(v reflect.Value) map[string]any {
	kwargs := make(map[string]any)

	// Get the type of the current struct
	t := v.Type()

	// Iterate through the fields and collect exported fields
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		if field.PkgPath == "" { // Only include exported fields
			kwargs[field.Name] = v.Field(i).Interface()
		}
	}

	return kwargs
}

// isValidIdentifier checks if a string is a valid Go identifier.
func isValidIdentifier(s string) bool {
	if len(s) == 0 {
		return false
	}
	for i, r := range s {
		if i == 0 && !isLetter(r) {
			return false
		}
		if i > 0 && !isLetterOrDigit(r) {
			return false
		}
	}
	return true
}

// isLetter checks if a rune is a valid Go letter.
func isLetter(r rune) bool {
	return ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z') || r == '_'
}

// isLetterOrDigit checks if a rune is a valid Go letter or digit.
func isLetterOrDigit(r rune) bool {
	return isLetter(r) || ('0' <= r && r <= '9')
}
