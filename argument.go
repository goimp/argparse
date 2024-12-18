package argparse

// Adding argument actions

type Argument struct {
	OptionStrings []string // The command-line option strings
	Dest          string   // The destination name where the value will be stored
	Nargs         any      // The number of arguments to consume
	Const         any      // The constant value for certain actions
	Default       any      // The default value if the option is not specified
	Type          Type     // The function to convert the string to the appropriate type
	Choices       []any    // The valid values for this argument
	Required      bool     // Whether the argument is required
	Help          string   // The help description for the argument
	MetaVar       any      // The name to be used in help output
	Deprecated    bool     // Whether the argument is deprecated
	Action        string
	Version       string
}
