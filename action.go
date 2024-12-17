package argparse

import (
	"argparse/attribute_holder" // Import the AttributeHolder package
)

type Type = any
type TypeFunc = func(string) (any, error)
type NewActionFuncType = func(*Argument) *Action

type ActionInterface interface {
	GetKwargs() map[string]any
	FormatUsage() string
	Call(parser *ArgumentParser, namespace *Namespace, values any, optionString string) error
	Self() *Action
}

// Action represents the action associated with an argument.
type Action struct {
	attribute_holder.AttributeHolder // Embedding AttributeHolder for its functionality

	OptionStrings []string // The command-line option strings
	Dest          string   // The destination name where the value will be stored
	Nargs         any      // The number of arguments to consume
	Const         any      // The constant value for certain actions
	Default       any      // The default value if the option is not specified
	Type          Type     // The function to convert the string to the appropriate type
	Choices       []any    // The valid values for this argument
	Required      bool     // Whether the argument is required
	Help          string   // The help description for the argument
	Metavar       any      // The name to be used in help output
	Deprecated    bool     // Whether the argument is deprecated

	container any
}

func NewAction(argument *Argument) *Action {
	return &Action{
		OptionStrings: argument.OptionStrings,
		Dest:          argument.Dest,
		Nargs:         argument.Nargs,
		Const:         argument.Const,
		Default:       argument.Default,
		Type:          argument.Type,
		Choices:       argument.Choices,
		Required:      argument.Required,
		Help:          argument.Help,
		Metavar:       argument.Metavar,
		Deprecated:    argument.Deprecated,
	}
}

func (a *Action) Self() *Action {
	return a
}

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
func (a *Action) Call(parser *ArgumentParser, namespace *Namespace, values any, optionString string) error {
	panic("action.Call() not implemented")
}
