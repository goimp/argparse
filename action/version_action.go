package action

import (
	"argparse/namespace"
	"fmt"
)

// VersionAction represents an action that displays the version information.
type VersionAction struct {
	Action
	Version string
}

// NewVersionAction creates a new VersionAction.
func NewVersionAction(
	optionStrings []string,
	version string,
	dest string,
	defaultVal any,
	help string,
	deprecated bool,
) (*VersionAction, error) {
	return &VersionAction{
		Action: Action{
			OptionStrings: optionStrings,
			Dest:          dest,
			Nargs:         0,
			Default:       defaultVal,
			Help:          help,
			Deprecated:    deprecated,
		},
		Version: version,
	}, nil
}

// SetValue prints the version information and exits the program.
func (a *VersionAction) Call(parser any, namespace *namespace.Namespace, values any, optionString string) {
	// Print the version information
	// version := a.Version
	fmt.Println("UNDONE PARSER VERSION")

	// if version == "" {
	// version = parser.Version // Default version if not specified
	// }
	// formatter := parser.GetFormatter()
	// formatter.AddText(version)
	// parser.PrintMessage(formatter.formatHelp(), os.Stdout)
	fmt.Printf("Version: %s\n", a.Version)

	// Exit the program after printing the version
	// os.Exit(0)
}
