package action

import (
	"fmt"
	"os"
)

// VersionAction represents an action that displays the version information.
type VersionAction struct {
	Action
	Version string
}

// NewVersionAction creates a new VersionAction.
func NewVersionAction(optionStrings []string, version string) *VersionAction {
	return &VersionAction{
		Action:  Action{OptionStrings: optionStrings},
		Version: version,
	}
}

// SetValue prints the version information and exits the program.
func (a *VersionAction) SetValue() {
	// Print the version information
	if a.Version == "" {
		a.Version = "1.0.0" // Default version if not specified
	}
	fmt.Printf("Version: %s\n", a.Version)

	// Exit the program after printing the version
	os.Exit(0)
}
