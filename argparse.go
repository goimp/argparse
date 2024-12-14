package argparse

import (
	"fmt"
)

const (
	SUPPRESS               = "==SUPPRESS=="
	OPTIONAL               = "?"
	ZERO_OR_MORE           = "*"
	ONE_OR_MORE            = "+"
	PARSER                 = "A..."
	REMAINDER              = "..."
	UNRECOGNIZED_ARGS_ATTR = "_unrecognized_args"
)

// Greet returns a greeting message.
func Greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
