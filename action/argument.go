package action

// Argument represents an argument with various properties similar to the Python version.
type Argument struct {
	OptionStrings []string      // Equivalent to `argument.option_strings`
	Metavar       interface{}   // Equivalent to `argument.metavar`, can be nil, string, or []string
	Nargs         int           // Equivalent to `argument.nargs`
	Dest          string        // Equivalent to `argument.dest`
	Choices       []interface{} // Equivalent to `argument.choices`
}

// Constants for nargs values
const (
	ZERO_OR_MORE = -1
	ONE_OR_MORE  = -2
	SUPPRESS     = ""
)
