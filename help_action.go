package argparse

// HelpAction represents an action that displays the help message.
type HelpAction struct {
	Action
}

// NewHelpAction creates a new HelpAction.
// NewHelpAction creates a new HelpAction.
func NewHelpAction(argument *Argument) ActionInterface {

	// Default dest to SUPPRESS if empty
	if argument.Dest == "" {
		argument.Dest = SUPPRESS
	}

	// Default defaultVal to SUPPRESS if nil
	if argument.Default == nil {
		argument.Default = SUPPRESS
	}

	// Create and return the HelpAction instance
	return &HelpAction{
		Action: Action{
			OptionStrings: argument.OptionStrings,
			Dest:          argument.Dest,
			Nargs:         0,
			Default:       argument.Default,
			Help:          argument.Help,
			Deprecated:    argument.Deprecated,
		},
	}
}

// SetValue prints the help message and exits the program.
func (a *HelpAction) Call(parser *ArgumentParser, namespace *Namespace, values any, optionString string) error {
	if parser != nil {
		parser.PrintHelp(nil)
		parser.Exit(0, "")
	}
	return nil
}
