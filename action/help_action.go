package action

import (
	"fmt"
	"os"
)

// HelpAction represents an action that displays the help message.
type HelpAction struct {
	Action
	OptionStrings []string
}

// NewHelpAction creates a new HelpAction.
func NewHelpAction(optionStrings []string) *HelpAction {
	return &HelpAction{
		Action:        Action{OptionStrings: optionStrings},
		OptionStrings: optionStrings,
	}
}

// SetValue prints the help message and exits the program.
func (a *HelpAction) SetValue() {
	// Print the help message (for demonstration purposes, we'll use a simple message)
	fmt.Println("Usage: [options]")
	fmt.Println("Options:")
	fmt.Println("  --help        Show this help message")

	// Exit the program after printing help
	os.Exit(0)
}
