package action

import (
	"fmt"
)

// Action represents the action associated with an argument.
type Action struct {
	OptionStrings []string                          // The command-line option strings
	Dest          string                            // The destination name where the value will be stored
	Nargs         *int                              // The number of arguments to consume
	Const         interface{}                       // The constant value for certain actions
	Default       interface{}                       // The default value if the option is not specified
	Type          func(string) (interface{}, error) // The function to convert the string to the appropriate type
	Choices       []interface{}                     // The valid values for this argument
	Required      bool                              // Whether the argument is required
	Help          string                            // The help description for the argument
	Metavar       string                            // The name to be used in help output
	Deprecated    bool                              // Whether the argument is deprecated
}

// NewAction creates a new Action object.
func NewAction(optionStrings []string, dest string, nargs *int, constVal interface{}, defaultVal interface{}, typ func(string) (interface{}, error), choices []interface{}, required bool, help, metavar string, deprecated bool) *Action {
	return &Action{
		OptionStrings: optionStrings,
		Dest:          dest,
		Nargs:         nargs,
		Const:         constVal,
		Default:       defaultVal,
		Type:          typ,
		Choices:       choices,
		Required:      required,
		Help:          help,
		Metavar:       metavar,
		Deprecated:    deprecated,
	}
}

// GetKwargs returns the keyword arguments associated with the Action.
func (a *Action) GetKwargs() map[string]interface{} {
	return map[string]interface{}{
		"optionStrings": a.OptionStrings,
		"dest":          a.Dest,
		"nargs":         a.Nargs,
		"const":         a.Const,
		"default":       a.Default,
		"type":          a.Type,
		"choices":       a.Choices,
		"required":      a.Required,
		"help":          a.Help,
		"metavar":       a.Metavar,
		"deprecated":    a.Deprecated,
	}
}

// FormatUsage returns the formatted usage for this action.
func (a *Action) FormatUsage() string {
	return a.OptionStrings[0]
}

// Call simulates the action being triggered (not implemented here, as per Python's version).
func (a *Action) Call(parser interface{}, namespace interface{}, values interface{}, optionString string) error {
	return fmt.Errorf("action .__call__() not implemented")
}

// // Example function to simulate the `Type` functionality (conversion of string to a specific type).
// func stringToInt(value string) (interface{}, error) {
// 	return fmt.Sprintf("Converted: %s", value), nil // Example conversion
// }

// // Example usage of Action.
// func main() {
// 	// Initialize Action
// 	nargs := 1
// 	action := NewAction([]string{"--verbose", "-v"}, "verbose", &nargs, nil, nil, stringToInt, nil, false, "Enable verbose output", "VERBOSE", false)

// 	// Print the formatted usage
// 	fmt.Println(action.FormatUsage()) // Output: --verbose

// 	// Print the kwargs
// 	for key, value := range action.GetKwargs() {
// 		fmt.Printf("%s: %v\n", key, value)
// 	}
// }
