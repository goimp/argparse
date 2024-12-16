package argparse

import (
	"argparse/namespace"
)

// HelpAction represents an action that displays the help message.
type HelpAction struct {
	Action
}

// NewHelpAction creates a new HelpAction.
// NewHelpAction creates a new HelpAction.
func NewHelpAction(
	optionStrings []string,
	dest string,
	defaultVal any,
	help string,
	deprecated bool,
) (*HelpAction, error) {

	// Default dest to SUPPRESS if empty
	if dest == "" {
		dest = SUPPRESS
	}

	// Default defaultVal to SUPPRESS if nil
	if defaultVal == nil {
		defaultVal = SUPPRESS
	}

	// Create and return the HelpAction instance
	return &HelpAction{
		Action: Action{
			OptionStrings: optionStrings,
			Dest:          dest,
			Nargs:         0,
			Default:       defaultVal,
			Help:          help,
			Deprecated:    deprecated,
		},
	}, nil
}

// SetValue prints the help message and exits the program.
func (a *HelpAction) Call(parser any, namespace *namespace.Namespace, values any, optionString string) {
	// if parser != nil {
	// 	parser.PrintHelp(nil)
	// 	parser.Exit(0, "")
	// }
}
