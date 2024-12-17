package argparse

import (
	"fmt"
	"os"
)

// VersionAction represents an action that displays the version information.
type VersionAction struct {
	*Action
	Version string
}

// NewVersionAction creates a new VersionAction.
func NewVersionAction(argument *Argument) *VersionAction {
	return &VersionAction{
		Action: &Action{
			OptionStrings: argument.OptionStrings,
			Dest:          argument.Dest,
			Nargs:         0,
			Default:       argument.Default,
			Help:          argument.Help,
			Deprecated:    argument.Deprecated,
		},
		Version: argument.Version,
	}
}

// FIXME: not done
// SetValue prints the version information and exits the program.
func (a *VersionAction) Call(parser *ArgumentParser, namespace *Namespace, values any, optionString string) error {
	if parser != nil {
		// Print the version information
		// version := a.Version
		fmt.Println("UNDONE PARSER VERSION")

		// if version == "" {
		// 	version = parser.Version // Default version if not specified
		// }
		// formatter := parser.GetFormatter()
		// formatter.AddText(version)
		// parser.PrintMessage(formatter.formatHelp(), os.Stdout)
		fmt.Printf("Version: %s\n", a.Version)

		// Exit the program after printing the version
		os.Exit(0)
	}
	return nil
}
