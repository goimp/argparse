package action

import (
	"argparse/attribute_holder" // Import the AttributeHolder package
	"argparse/namespace"        // Import the AttributeHolder package
)

type StringToAnyFunc func(string) (any, error)

// Action represents the action associated with an argument.
type Action struct {
	attribute_holder.AttributeHolder // Embedding AttributeHolder for its functionality

	OptionStrings []string        // The command-line option strings
	Dest          string          // The destination name where the value will be stored
	Nargs         any             // The number of arguments to consume
	Const         any             // The constant value for certain actions
	Default       any             // The default value if the option is not specified
	Type          StringToAnyFunc // The function to convert the string to the appropriate type
	Choices       []any           // The valid values for this argument
	Required      bool            // Whether the argument is required
	Help          string          // The help description for the argument
	Metavar       string          // The name to be used in help output
	Deprecated    bool            // Whether the argument is deprecated
}

func NewAction(
	optionStrings []string,
	dest string,
	nargs any,
	constVal any,
	defaultVal any,
	typ StringToAnyFunc,
	choices []any,
	required bool,
	help string,
	metavar string,
	deprecated bool,
) (*Action, error) {
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
	}, nil
}

// // Override GetArgs if needed
// func (a *Action) GetArgs() []string {
// 	// Example: No positional arguments for now
// 	return []string{}
// }

// Override GetKwargs to customize keyword arguments
func (a *Action) GetKwargs() map[string]any {
	return map[string]any{
		"OptionStrings": a.OptionStrings,
		"Dest":          a.Dest,
		"Nargs":         a.Nargs,
		"Const":         a.Const,
		"Default":       a.Default,
		"Type":          a.Type,
		"Choices":       a.Choices,
		"Required":      a.Required,
		"Help":          a.Help,
		"Metavar":       a.Metavar,
		"Deprecated":    a.Deprecated,
	}
}

// FormatUsage returns the formatted usage for this action.
func (a *Action) FormatUsage() string {
	if len(a.OptionStrings) > 0 {
		return a.OptionStrings[0]
	}
	return ""
}

// Call simulates the action being triggered (not implemented here, as per Python's version).
func (a *Action) Call(parser any, namespace *namespace.Namespace, values any, optionString string) error {
	panic("action.Call() not implemented")
}
