package action

import (
	"fmt"
	"reflect"
	"strings"
)

// BooleanOptionalAction represents a boolean flag action with support for both `--flag` and `--no-flag`.
type BooleanOptionalAction struct {
	*Action // Inherit from Action

	// No additional fields are needed as the Action struct already includes everything
}

// NewBooleanOptionalAction creates a new BooleanOptionalAction object.
func NewBooleanOptionalAction(optionStrings []string, dest string, defaultVal interface{}, required bool, help string, deprecated bool) (*BooleanOptionalAction, error) {
	// Validate and process option strings
	var _optionStrings []string
	for _, optionString := range optionStrings {
		_optionStrings = append(_optionStrings, optionString)

		// Check for invalid --no- prefix
		if strings.HasPrefix(optionString, "--no-") {
			return nil, fmt.Errorf("invalid option name %q for BooleanOptionalAction", optionString)
		}

		// Add the --no- prefixed version of the option
		noOption := "--no-" + optionString[2:]
		_optionStrings = append(_optionStrings, noOption)
	}

	// Create Action (parent struct) with nargs=0 for BooleanOptionalAction
	action := &Action{
		OptionStrings: _optionStrings,
		Dest:          dest,
		Nargs:         new(int), // Set to 0 to indicate no arguments needed
		Default:       defaultVal,
		Required:      required,
		Help:          help,
		Metavar:       "",
		Deprecated:    deprecated,
	}

	return &BooleanOptionalAction{Action: action}, nil
}

// Call executes the action when the option is encountered on the command line.
func (a *BooleanOptionalAction) Call(parser interface{}, namespace interface{}, values interface{}, optionString string) error {
	// Check if the option string is in the list of option strings
	if contains(a.OptionStrings, optionString) {
		// Toggle the value: if option starts with --no-, set false, else set true
		value := true
		if strings.HasPrefix(optionString, "--no-") {
			value = false
		}

		// Set the value in the namespace (assuming it's a pointer to the destination)
		setValue(namespace, a.Dest, value)
	}

	return nil
}

// FormatUsage formats the usage for the action, combining the option strings with a separator.
func (a *BooleanOptionalAction) FormatUsage() string {
	return strings.Join(a.OptionStrings, " | ")
}

// Helper function to check if a slice contains a string
func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

// Helper function to set a value in the namespace (this is a placeholder function)
// setValue sets the value in the namespace (assuming it's a pointer to a struct).
func setValue(namespace interface{}, dest string, value interface{}) error {
	// Ensure that the namespace is a pointer to a struct
	v := reflect.ValueOf(namespace)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("namespace must be a pointer to a struct")
	}

	// Get the field by name (dest should match the struct field name)
	field := v.Elem().FieldByName(dest)
	if !field.IsValid() {
		return fmt.Errorf("no such field: %s in struct", dest)
	}

	// Ensure that the field can be set (it must be exported)
	if !field.CanSet() {
		return fmt.Errorf("cannot set field: %s", dest)
	}

	// Set the field value
	field.Set(reflect.ValueOf(value))

	return nil
}

// func main() {
// 	// Example usage of BooleanOptionalAction

// 	// Initialize the action
// 	action, err := NewBooleanOptionalAction([]string{"--verbose"}, "verbose", nil, false, "Enable verbose output", false)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	// Print the usage
// 	fmt.Println(action.FormatUsage()) // Output: --verbose | --no-verbose

// 	// Simulate calling the action when a flag is encountered
// 	// (In reality, you'd pass a parser, namespace, etc.)
// 	action.Call(nil, nil, nil, "--verbose") // Output: Setting value for verbose: true
// 	action.Call(nil, nil, nil, "--no-verbose") // Output: Setting value for verbose: false
// }
