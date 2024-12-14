package action

import (
	"fmt"
	"strings"
)

// SubParsersAction handles subcommands and their parsers
type SubParsersAction struct {
	// A map of subparser names to their respective SubParser
	NameParserMap map[string]*SubParser
	// A list of all defined subparser actions for help and choices
	ChoicesActions []*SubParser
	// A set of deprecated subparser names
	Deprecated map[string]bool
	// The program prefix (for constructing subcommand help)
	ProgPrefix string
}

// NewSubParsersAction creates a new SubParsersAction instance
func NewSubParsersAction(progPrefix string) *SubParsersAction {
	return &SubParsersAction{
		NameParserMap:  make(map[string]*SubParser),
		ChoicesActions: []*SubParser{},
		Deprecated:     make(map[string]bool),
		ProgPrefix:     progPrefix,
	}
}

// AddParser adds a new subparser to the action
func (spa *SubParsersAction) AddParser(name string, handler func(args []string) error, aliases []string, help string, deprecated bool) (*SubParser, error) {
	// Check for conflicts with existing subparser names or aliases
	if _, exists := spa.NameParserMap[name]; exists {
		return nil, fmt.Errorf("conflicting subparser: %s", name)
	}

	for _, alias := range aliases {
		if _, exists := spa.NameParserMap[alias]; exists {
			return nil, fmt.Errorf("conflicting subparser alias: %s", alias)
		}
	}

	// Create the new subparser and add it to the map
	subparser := &SubParser{
		Name:    name,
		Aliases: aliases,
		Help:    help,
		Handler: handler,
	}

	spa.NameParserMap[name] = subparser
	spa.ChoicesActions = append(spa.ChoicesActions, subparser)

	// Make the subparser available under aliases as well
	for _, alias := range aliases {
		spa.NameParserMap[alias] = subparser
	}

	// Mark as deprecated if necessary
	if deprecated {
		spa.Deprecated[name] = true
		for _, alias := range aliases {
			spa.Deprecated[alias] = true
		}
	}

	return subparser, nil
}

// HandleSubcommand processes the subcommand and delegates to the appropriate handler
func (spa *SubParsersAction) HandleSubcommand(subcommand string, args []string) error {
	// Select the subparser based on the subcommand name
	subparser, exists := spa.NameParserMap[subcommand]
	if !exists {
		return fmt.Errorf("unknown subcommand %q (choices: %s)", subcommand, spa.getChoices())
	}

	// If the subparser is deprecated, print a warning
	if spa.Deprecated[subcommand] {
		fmt.Printf("Warning: The subcommand '%s' is deprecated\n", subcommand)
	}

	// Call the handler for the subparser
	return subparser.Handler(args)
}

// getChoices returns a string of available subcommands and aliases
func (spa *SubParsersAction) getChoices() string {
	var choices []string
	for name := range spa.NameParserMap {
		choices = append(choices, name)
	}
	return strings.Join(choices, ", ")
}

// _ExtendAction is an action that extends an existing list with new values
type _ExtendAction struct {
	Action
}

// NewExtendAction creates a new _ExtendAction
func NewExtendAction(optionStrings []string, dest string) *_ExtendAction {
	return &_ExtendAction{
		Action: Action{
			OptionStrings: optionStrings,
			Dest:          dest,
		},
	}
}

// SetValue extends the list in the destination with new values
func (a *_ExtendAction) SetValue(namespace map[string]interface{}, values []interface{}) {
	// Get the current value from the namespace
	items, exists := namespace[a.Dest]
	if !exists {
		// If the items don't exist, initialize an empty slice
		items = []interface{}{}
	}

	// Ensure items is a slice
	itemSlice, ok := items.([]interface{})
	if !ok {
		// If not a slice, initialize a new slice
		itemSlice = []interface{}{items}
	}

	// Extend the slice with new values
	itemSlice = append(itemSlice, values...)

	// Set the extended list back in the namespace
	namespace[a.Dest] = itemSlice
}
