package argparse

// SubParser represents a subcommand parser in the CLI
type SubParser struct {
	// The name of the subparser
	Name string
	// Aliases for the subparser (alternative names)
	Aliases []string
	// The help message for the subparser
	Help string
	// The function to handle this subparser's arguments
	Handler func(args []string) error
}
