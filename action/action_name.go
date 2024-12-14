package action

import (
	"fmt"
	"strings"
)

// GetActionName determines the action name for the given argument.
func GetActionName(argument *Argument) string {
	if argument == nil {
		return ""
	}

	// If option strings are present
	if len(argument.OptionStrings) > 0 {
		return strings.Join(argument.OptionStrings, "/")
	}

	// If metavar is present
	if metavar, ok := argument.Metavar.(string); ok && metavar != SUPPRESS {
		return metavar
	} else if metavar, ok := argument.Metavar.([]string); ok {
		if argument.Nargs == ZERO_OR_MORE && len(metavar) == 2 {
			return fmt.Sprintf("%s[, %s]", metavar[0], metavar[1])
		} else if argument.Nargs == ONE_OR_MORE {
			return fmt.Sprintf("%s[, %s]", metavar[0], metavar[1])
		}
		return strings.Join(metavar, ", ")
	}

	// If destination is present
	if argument.Dest != "" && argument.Dest != SUPPRESS {
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

// func main() {
// 	// Example usage

// 	// Case with option strings
// 	arg1 := &Argument{
// 		OptionStrings: []string{"--verbose", "-v"},
// 	}
// 	fmt.Println(GetActionName(arg1)) // Output: --verbose/-v

// 	// Case with metavar (single)
// 	arg2 := &Argument{
// 		Metavar: "FILE",
// 	}
// 	fmt.Println(GetActionName(arg2)) // Output: FILE

// 	// Case with metavar (tuple)
// 	arg3 := &Argument{
// 		Metavar: []string{"SRC", "DEST"},
// 		Nargs:   ZERO_OR_MORE,
// 	}
// 	fmt.Println(GetActionName(arg3)) // Output: SRC[, DEST]

// 	// Case with destination
// 	arg4 := &Argument{
// 		Dest: "output",
// 	}
// 	fmt.Println(GetActionName(arg4)) // Output: output

// 	// Case with choices
// 	arg5 := &Argument{
// 		Choices: []interface{}{"red", "green", "blue"},
// 	}
// 	fmt.Println(GetActionName(arg5)) // Output: {red,green,blue}

// 	// Default case
// 	arg6 := &Argument{}
// 	fmt.Println(GetActionName(arg6)) // Output: ""
// }
