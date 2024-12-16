package action

import (
	"argparse"
	"fmt"
	"strings"
)

// GetActionName determines the action name for the given argument.
func GetActionName(argument *Action) string {
	if argument == nil {
		return ""
	}

	// If option strings are present
	if len(argument.OptionStrings) > 0 {
		return strings.Join(argument.OptionStrings, "/")
	}

	// If metavar is present
	if metavar, ok := argument.Metavar.(string); ok && metavar != argparse.SUPPRESS {
		return metavar
	} else if metavar, ok := argument.Metavar.([]string); ok {
		if argument.Nargs == argparse.ZERO_OR_MORE && len(metavar) == 2 {
			return fmt.Sprintf("%s[, %s]", metavar[0], metavar[1])
		} else if argument.Nargs == argparse.ONE_OR_MORE {
			return fmt.Sprintf("%s[, %s]", metavar[0], metavar[1])
		}
		return strings.Join(metavar, ", ")
	}

	// If destination is present
	if argument.Dest != "" && argument.Dest != argparse.SUPPRESS {
		return argument.Dest
	}

	// If choices are present
	if len(argument.Choices) > 0 {
		choices := make([]string, len(argument.Choices))
		for i, choice := range argument.Choices {
			choices[i] = fmt.Sprintf("%v", choice)
		}
		return fmt.Sprintf("{%s}", strings.Join(choices, ","))
	}

	// Default case
	return ""
}
