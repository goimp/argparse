package argparse

import (
	"fmt"
	"regexp"
)

type ActionsContainer struct {
	Description                string
	PrefixChars                any
	ArgumentDefault            any
	ConflictHandler            any
	registries                 map[any]any
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
	prefixChars any,
	argumentDefault any,
	conflictHandler any,
) (*ActionsContainer, error) {

	container := &ActionsContainer{
		Description:                description,
		PrefixChars:                prefixChars,
		ArgumentDefault:            argumentDefault,
		ConflictHandler:            conflictHandler,
		registries:                 make(map[any]any), // set up registries
		actions:                    []Action{},        // action storage
		optionStringActions:        make(map[any]any),
		actionGroups:               []any{}, // groups
		mutuallyExclusiveGroups:    []any{},
		defaults:                   make(map[string]any),            // defaults storage
		negativeNumberMatcher:      regexp.MustCompile(`-\.?(\d+)`), // determines whether an "option" looks like a negative number
		hasNegativeNumberOptionals: []any{},                         // # whether or not there are any optionals that look like negative numbers -- uses a list so it can be shared and edited
	}

	// register actions
	container.Register("action", nil, NewStoreAction)
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

func (a *ActionsContainer) Register(registryName string, value any, object any) {
	// Check if the registry exists
	registry, exists := a.registries[registryName]

	if !exists {
		// Initialize the registry as a map[any]any
		registry = make(map[any]any)
		a.registries[registryName] = registry
	}

	// Perform a type assertion to ensure registry is map[any]any
	if regMap, ok := registry.(map[any]any); ok {
		// Add the value-object pair to the registry
		regMap[value] = object
	} else {
		panic(fmt.Sprintf("registry %s is not of type map[any]any", registryName))
	}
}

func (a *ActionsContainer) RegistryGet(registryName string, value any, defaultValue any) any {
	// Check if the registry exists
	registry, exists := a.registries[registryName]
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
func (a *ActionsContainer) SetDefaults(kwargs map[string]any) {
	// Update the _defaults map
	for key, value := range kwargs {
		a.defaults[key] = value
	}

	// Update the default value of actions that match the keys in kwargs
	for i := range a.actions {
		if defaultValue, exists := kwargs[a.actions[i].Dest]; exists {
			a.actions[i].Default = defaultValue
		}
	}
}

func (a *ActionsContainer) GetDefault(dest string) any {
	// Iterate over actions to find a matching dest with a non-nil default value
	for _, action := range a.actions {
		if action.Dest == dest && action.Default != nil {
			return action.Default
		}
	}
	// Return the default from the _defaults map if it exists
	return a.defaults[dest]
}

// Adding argument actions

func (a *ActionsContainer) AddArgument(args []any, kwargs map[string]any) {

}

func (a *ActionsContainer) AddArgumentGroup(args []any, kwargs map[string]any) {

}

func (a *ActionsContainer) AddMutuallyExclusiveGroup(kwargs map[string]any) {

}

func (a *ActionsContainer) AddAction(action *Action) {

}

func (a *ActionsContainer) RemoveAction(action *Action) {

}

func (a *ActionsContainer) AddContainerAction(container *ActionsContainer) {

}

func (a *ActionsContainer) GetPositionalKwargs(dest string, kwargs map[string]any) {

}

func (a *ActionsContainer) GetOptionalKwargs(args []any, kwargs map[string]any) {

}

func (a *ActionsContainer) PopActionClass(kwargs map[string]any, defaultAction *Action) {

}

func (a *ActionsContainer) GetHandler() (func(), error) {
	switch a.ConflictHandler {
	case "ignore":
		return func() {
			fmt.Println("Ignoring conflicts")
		}, nil
	case "resolve":
		return func() {
			fmt.Println("Resolving conflicts")
		}, nil
	default:
		return nil, fmt.Errorf("invalid conflict resolution value: %v", a.ConflictHandler)
	}
}

func (a *ActionsContainer) CheckConflict(action *Action) {

}

func (a *ActionsContainer) HandleConflictError(action *Action, conflictingActions []*Action) {

}

func (a *ActionsContainer) HandleConflictResolve(action *Action, conflictingActions []*Action) {

}

func (a *ActionsContainer) CheckHelp(action *Action) {
	// if action.Help
}
