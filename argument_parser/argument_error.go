package argument_parser

import (
	"argparse/action"
	"fmt"
)

// ArgumentError represents an error that occurs during argument creation or usage.
type ArgumentError struct {
	ArgumentName string
	Message      string
}

// NewArgumentError creates a new ArgumentError with a formatted message.
func NewArgumentError(argument *action.Action, message string) *ArgumentError {
	return &ArgumentError{
		ArgumentName: action.GetActionName(argument),
		Message:      message,
	}
}

// Error implements the error interface for ArgumentError.
func (e *ArgumentError) Error() string {
	if e.ArgumentName == "" {
		return e.Message
	}
	return fmt.Sprintf("argument %s: %s", e.ArgumentName, e.Message)
}

// ArgumentTypeError represents an error when trying to convert a command-line string to a specific type.
type ArgumentTypeError struct {
	Message string
}

// NewArgumentTypeError creates a new ArgumentTypeError.
func NewArgumentTypeError(message string) *ArgumentTypeError {
	return &ArgumentTypeError{
		Message: message,
	}
}

// Error implements the error interface for ArgumentTypeError.
func (e *ArgumentTypeError) Error() string {
	return e.Message
}

// // Example usage of ArgumentError within the argparse package.
// func main() {
// 	// Example with valid argument details
// 	arg := &Argument{
// 		OptionStrings: []string{"--verbose"},
// 		Dest:          "verbose",
// 	}

// 	err := NewArgumentError(arg, "invalid value")
// 	if err != nil {
// 		fmt.Println(err.Error()) // Output: argument --verbose: invalid value
// 	}

// 	// Example with no argument name
// 	err = NewArgumentError(nil, "missing argument")
// 	fmt.Println(err.Error()) // Output: missing argument

// 	// Example with metavar
// 	arg2 := &Argument{
// 		Metavar: "FILE",
// 	}
// 	err = NewArgumentError(arg2, "file not found")
// 	fmt.Println(err.Error()) // Output: argument FILE: file not found
// }

// // Example usage of ArgumentTypeError.
// func main() {
// 	// Example with a type conversion error
// 	err := NewArgumentTypeError("invalid type for argument")
// 	if err != nil {
// 		fmt.Println(err.Error()) // Output: invalid type for argument
// 	}
// }
