package argparse

import (
	"fmt"
	"regexp"
)

type ActionsContainer struct {
	Description                string
	PrefixChars                []string
	ArgumentDefault            any
	ConflictHandler            any
	registries                 map[string]any
	actions                    []Action
	optionStringActions        map[any]any
	actionGroups               []any
	mutuallyExclusiveGroups    []any
	defaults                   map[string]any
	negativeNumberMatcher      *regexp.Regexp
	hasNegativeNumberOptionals []any
}

func NewActionsContainer(
	description string,
	prefixChars []string,
	argumentDefault any,
	conflictHandler any,
) (*ActionsContainer, error) {

	container := &ActionsContainer{
		Description:                description,
		PrefixChars:                prefixChars,
		ArgumentDefault:            argumentDefault,
		ConflictHandler:            conflictHandler,
		registries:                 make(map[string]any), // set up registries
		actions:                    []Action{},           // action storage
		optionStringActions:        make(map[any]any),
		actionGroups:               []any{}, // groups
		mutuallyExclusiveGroups:    []any{},
		defaults:                   make(map[string]any),            // defaults storage
		negativeNumberMatcher:      regexp.MustCompile(`-\.?(\d+)`), // determines whether an "option" looks like a negative number
		hasNegativeNumberOptionals: []any{},                         // # whether or not there are any optionals that look like negative numbers -- uses a list so it can be shared and edited
	}

	// register actions
	container.Register("action", "", NewStoreAction)
	container.Register("action", "store", NewStoreAction)
	container.Register("action", "store_const", NewStoreConstAction)
	container.Register("action", "store_true", NewStoreTrueAction)
	container.Register("action", "store_false", NewStoreFalseAction)
	container.Register("action", "append", NewAppendAction)
	container.Register("action", "append_const", NewAppendConstAction)
	container.Register("action", "count", NewCountAction)
	container.Register("action", "help", NewHelpAction)
	container.Register("action", "version", NewVersionAction)
	container.Register("action", "parsers", NewSubParsersAction)
	container.Register("action", "extend", NewExtendAction)

	// raise an exception if the conflict handler is invalid
	container.GetHandler()

	// action storage

	return container, nil
}

// Registration methods

func (ac *ActionsContainer) Register(registryName string, value string, object any) {
	// Check if the registry exists
	registry, exists := ac.registries[registryName]

	if !exists {
		// Initialize the registry as a map[any]any
		registry = make(map[string]any)
		ac.registries[registryName] = registry
	}

	// Perform a type assertion to ensure registry is map[any]any
	if regMap, ok := registry.(map[string]any); ok {
		// Add the value-object pair to the registry
		regMap[value] = object
	} else {
		panic(fmt.Sprintf("registry %s is not of type map[any]any", registryName))
	}
}

func (ac *ActionsContainer) RegistryGet(registryName string, value any, defaultValue any) any {
	// Check if the registry exists
	registry, exists := ac.registries[registryName]
	if !exists {
		return defaultValue // Return the default value if the registry doesn't exist
	}

	// Perform a type assertion to ensure registry is map[any]any
	if regMap, ok := registry.(map[any]any); ok {
		// Try to get the value from the registry
		if obj, found := regMap[value]; found {
			return obj // Return the value if found
		}
	}

	// Return the default value if the key doesn't exist or the registry is invalid
	return defaultValue
}

// Namespace default accessor methods

// SetDefaults updates the default values and the action defaults
func (ac *ActionsContainer) SetDefaults(kwargs map[string]any) {
	// Update the _defaults map
	for key, value := range kwargs {
		ac.defaults[key] = value
	}

	// Update the default value of actions that match the keys in kwargs
	for i := range ac.actions {
		if defaultValue, exists := kwargs[ac.actions[i].Dest]; exists {
			ac.actions[i].Default = defaultValue
		}
	}
}

func (ac *ActionsContainer) GetDefault(dest string) any {
	// Iterate over actions to find a matching dest with a non-nil default value
	for _, action := range ac.actions {
		if action.Dest == dest && action.Default != nil {
			return action.Default
		}
	}
	// Return the default from the _defaults map if it exists
	return ac.defaults[dest]
}

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
	Metavar       any      // The name to be used in help output
	Deprecated    bool     // Whether the argument is deprecated
}

func (ac *ActionsContainer) AddArgument(argument Argument) (any, error) {
	// chars := ac.PrefixChars

	// if argument.OptionStrings != nil || len(argument.OptionStrings) == 1 {

	// }

	// kwargs := ac.GetPositionalKwargs(argument)
	return nil, nil
}

func (ac *ActionsContainer) AddArgumentGroup(args []any, kwargs map[string]any) {
	// group := NewArgumentGroup(
	// 	a,
	// 	kwargs["title"],
	// 	kwargs["description"],

	// )
}

func (ac *ActionsContainer) AddMutuallyExclusiveGroup(kwargs map[string]any) {

}

func (ac *ActionsContainer) AddAction(action *Action) *Action {
	return &Action{}
}

func (ac *ActionsContainer) RemoveAction(action *Action) {

}

func (ac *ActionsContainer) AddContainerAction(container *ActionsContainer) {

}

func (ac *ActionsContainer) GetPositionalKwargs(argument *Argument) {

}

func (ac *ActionsContainer) GetOptionalKwargs(args []any, kwargs map[string]any) {

}

func (ac *ActionsContainer) PopActionClass(kwargs map[string]any, defaultAction *Action) {

}

func (ac *ActionsContainer) GetHandler() (func(), error) {
	switch ac.ConflictHandler {
	case "ignore":
		return func() {
			fmt.Println("Ignoring conflicts")
		}, nil
	case "resolve":
		return func() {
			fmt.Println("Resolving conflicts")
		}, nil
	default:
		return nil, fmt.Errorf("invalid conflict resolution value: %v", ac.ConflictHandler)
	}
}

func (ac *ActionsContainer) CheckConflict(action *Action) {

}

func (ac *ActionsContainer) HandleConflictError(action *Action, conflictingActions []*Action) {

}

func (ac *ActionsContainer) HandleConflictResolve(action *Action, conflictingActions []*Action) {

}

func (ac *ActionsContainer) CheckHelp(action *Action) {
	// if action.Help
}
