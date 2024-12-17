package argparse

import (
	"argparse/namespace"
	"fmt"
	"strings"
)

// BooleanOptionalAction represents a boolean flag action with support for both `--flag` and `--no-flag`.
type BooleanOptionalAction struct {
	Action // Embed Action to inherit its behavior
}

// NewBooleanOptionalAction creates a new BooleanOptionalAction object.
func NewBooleanOptionalAction(optionStrings []string, dest string, defaultVal any, required bool, help string, deprecated bool) *BooleanOptionalAction {
	// Validate and process option strings
	var _optionStrings []string

	for _, optionString := range optionStrings {
		_optionStrings = append(_optionStrings, optionString)

		if strings.HasPrefix(optionString, "--") {
			if strings.HasPrefix(optionString, "--no-") {
				panic(fmt.Errorf("invalid option name %q for BooleanOptionalAction", optionString))
			}
			optionString = "--no-" + optionString[2:]
			_optionStrings = append(_optionStrings, optionString)
		}
	}

	action := Action{
		OptionStrings: _optionStrings,
		Dest:          dest,
		Nargs:         0,
		Default:       defaultVal,
		Required:      required,
		Help:          help,
		Deprecated:    deprecated,
	}

	return &BooleanOptionalAction{Action: action}
}

// Call executes the action when the option is encountered on the command line.
func (a *BooleanOptionalAction) Call(parser *ArgumentParser, namespace *namespace.Namespace, values any, optionString string) error {
	for _, _optionString := range a.OptionStrings {
		if _optionString == optionString {
			namespace.Set(a.Dest, !strings.HasPrefix(_optionString, "--no-"))
			break
		}
	}
	return nil
}

// FormatUsage formats the usage for the action, combining the option strings with a separator.
func (a *BooleanOptionalAction) FormatUsage() string {
	return strings.Join(a.OptionStrings, " | ")
}
