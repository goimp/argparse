package argparse

import (
	"fmt"
	"regexp"
	"strings"
)

type Registry = map[any]any
type Registries = map[string]Registry

type ConflictingOption struct {
	OptionString   string
	ConflictAction ActionInterface
}

type ActionsContainerInterface interface {
	Struct() *ActionsContainer                                                     // +
	Register(string, any, any)                                                     // +
	RegistryGet(string, any, any) any                                              // +
	AddArgument(*Argument) ActionInterface                                         // ?
	SetDefaults(map[string]any)                                                    // +
	GetDefault(string) any                                                         // +
	AddArgumentGroup(ActionsContainerInterface) ActionsContainerInterface          // ?
	AddMutuallyExclusiveGroup(ActionsContainerInterface) ActionsContainerInterface // ?
	AddAction(ActionInterface) ActionInterface                                     // ?
	RemoveAction(ActionInterface)                                                  // ?
	AddContainerAction(ActionsContainerInterface)                                  // ?
	GetPositionalArgument(*Argument) *Argument                                     // +
	GetOptionalArgument(*Argument) *Argument                                       // +
	GetHandler() func(ActionInterface, []ConflictingOption)                        // -
	CheckConflict(ActionInterface)                                                 // ?
	HandleConflictError(ActionInterface, []ConflictingOption)                      // ?
	HandleConflictResolve(ActionInterface, []ConflictingOption)                    // ?
	CheckHelp(action ActionInterface)                                              // ?
}

type ActionsContainer struct {
	Description                string
	PrefixChars                string
	ArgumentDefault            any
	ConflictHandler            any
	Registries                 Registries
	Actions                    []ActionInterface
	OptionStringActions        map[string]ActionInterface
	ActionGroups               []ActionsContainerInterface
	MutuallyExclusiveGroups    []ActionsContainerInterface
	Defaults                   map[string]any
	NegativeNumberMatcher      *regexp.Regexp
	HasNegativeNumberOptionals []bool

	GetFormatter any
	Title        string
	Required     bool
}

func (ac *ActionsContainer) Struct() *ActionsContainer {
	return ac
}

func NewActionsContainer(
	description string,
	prefixChars string,
	argumentDefault any,
	conflictHandler any,
) ActionsContainerInterface {

	container := &ActionsContainer{
		Description:                description,
		PrefixChars:                prefixChars,
		ArgumentDefault:            argumentDefault,
		ConflictHandler:            conflictHandler,
		Registries:                 make(Registries),    // set up registries
		Actions:                    []ActionInterface{}, // action storage
		OptionStringActions:        make(map[string]ActionInterface),
		ActionGroups:               []ActionsContainerInterface{}, // groups
		MutuallyExclusiveGroups:    []ActionsContainerInterface{},
		Defaults:                   make(map[string]any),            // defaults storage
		NegativeNumberMatcher:      regexp.MustCompile(`-\.?(\d+)`), // determines whether an "option" looks like a negative number
		HasNegativeNumberOptionals: []bool{},                        // # whether or not there are any optionals that look like negative numbers -- uses a list so it can be shared and edited
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

	return container
}

// Registration methods

// Register method to add a value to the registry in ActionsContainer
func (ac *ActionsContainer) Register(registryName string, value any, object any) {
	// Set default registry if it doesn't exist
	registry, exists := ac.Registries[registryName]
	if !exists {
		registry = make(Registry)
		ac.Registries[registryName] = registry
	}

	// Add the value-object pair to the registry
	registry[value] = object
}

// _RegistryGet method to retrieve a value from a registry with a default
func (ac *ActionsContainer) RegistryGet(registryName string, value any, defaultVal any) any {
	// Retrieve the registry by name
	registry, exists := ac.Registries[registryName]
	if !exists {
		return defaultVal // Return default value if registry doesn't exist
	}

	// Retrieve the value from the registry
	if result, found := registry[value]; found {
		return result // Return the value if found
	}

	// Return the default value if the value is not found
	return defaultVal
}

// Namespace default accessor methods

// SetDefaults updates the default values and the action defaults
func (ac *ActionsContainer) SetDefaults(mapping map[string]any) {
	// Update the _defaults map
	for key, value := range mapping {
		ac.Defaults[key] = value
	}

	// Update the default value of actions that match the keys in mapping
	for _, actionInterface := range ac.Actions {
		action := actionInterface.Struct()
		if defaultValue, exists := mapping[action.Dest]; exists {
			action.Default = defaultValue
		}
	}
}

func (ac *ActionsContainer) GetDefault(dest string) any {
	// Iterate over actions to find a matching dest with a non-nil default value
	for _, actionInterface := range ac.Actions {
		action := actionInterface.Struct()
		if action.Dest == dest && action.Default != nil {
			return action.Default
		}
	}
	// Return the default from the _defaults map if it exists
	return ac.Defaults[dest]
}

func isArgInChars(value string, chars string) bool {
	for _, c := range chars {
		if string(c) == value {
			return true
		}
	}
	return false
}

func (ac *ActionsContainer) AddArgument(argument *Argument) ActionInterface {
	chars := ac.PrefixChars
	args := argument.OptionStrings

	// Check if args are empty or not of a valid prefix
	if args == nil || (len(args) == 1 && !isArgInChars(args[0][:1], chars)) {
		// Handle positional arguments
		if argument.Dest != "" {
			panic("Dest supplied twice for positional argument, did you mean MetaVar")
		}
		argument = ac.GetPositionalArgument(argument)
	} else {
		// Handle optional arguments
		argument = ac.GetOptionalArgument(argument)
	}

	// if no default was supplied, use the parser-level default
	if argument.Default == nil {
		dest := argument.Dest
		if _, exist := ac.Defaults[dest]; exist {
			argument.Default = ac.Defaults[dest]
		} else if ac.ArgumentDefault != nil {
			argument.Default = ac.ArgumentDefault
		}
	}

	// Check if the registry value is of type NewActionFuncType and invoke it
	actionName := argument.Action

	// action := ac.createAction(actionName, argument)

	createAction := ac.RegistryGet("action", actionName, actionName)
	callback, ok := createAction.(func(*Argument) ActionInterface)
	if !ok {
		panic(fmt.Sprintf("action %s, result does not implement ActionInterface: %T", actionName, callback))
	}
	action := callback(argument)

	// raise an error if action for positional argument does not consume arguments
	if action.Struct().OptionStrings == nil {
		if nargs, ok := action.Struct().Nargs.(int); ok && nargs == 0 {
			panic(fmt.Sprintf("action %v is not valid for positional arguments", actionName))
		}
	}

	// FIXME: not done
	// // raise an error if the action type is not callable
	// var actionValueType = ac.registryGet("type", action.Type, action.Type)

	// if typeFunc, ok := actionValueType.(TypeFunc); !ok {
	// 	panic(fmt.Sprintf("%v is not TypeFunc", typeFunc))
	// }

	// FIXME: not done
	// if _, ok = actionValueType.(FileType); ok {
	// 	panic(fmt.Sprintf("%v is a FileType, instance of it must be passed", typeFunc))
	// }

	// FIXME: not done
	// raise an error if the metavar does not match the type
	// if ac.GetFormatter != nil {
	// 	formatter := ac.GetFormatter.(FormatterFunc)()
	// 	if err := formatter.FormatArgs(action, nil); err {
	// 		panic("length of metavar tuple does not match nargs")
	// 	}
	// }
	ac.CheckHelp(action)

	return ac.AddAction(action)
}

// FIXME: not done
func (ac *ActionsContainer) AddArgumentGroup(argumentGroup ActionsContainerInterface) ActionsContainerInterface {
	// group := NewArgumentGroup(
	// 	a,
	// 	kwargs["title"],
	// 	kwargs["description"],

	// )
	return argumentGroup
}

func (ac *ActionsContainer) AddMutuallyExclusiveGroup(mutuallyExclusiveGroup ActionsContainerInterface) ActionsContainerInterface {
	return mutuallyExclusiveGroup
}

func (ac *ActionsContainer) AddAction(action ActionInterface) ActionInterface {
	// resolve any conflicts
	ac.CheckConflict(action)

	// add to actions list
	ac.Actions = append(ac.Actions, action)
	action.Struct().Container = ac

	// index the action by any option strings it has
	for _, optionString := range action.Struct().OptionStrings {
		ac.OptionStringActions[optionString] = action
	}

	// set the flag if any option strings look like negative numbers
	for _, optionString := range action.Struct().OptionStrings {
		if ac.NegativeNumberMatcher.MatchString(optionString) {
			if len(ac.HasNegativeNumberOptionals) == 0 {
				ac.HasNegativeNumberOptionals = append(ac.HasNegativeNumberOptionals, true)
			}
		}
	}
	// return the created action
	return action
}

func (ac *ActionsContainer) RemoveAction(action ActionInterface) {
	for i, v := range ac.Actions {
		if v == action {
			// Remove the item by slicing the array
			ac.Actions = append(ac.Actions[:i], ac.Actions[i+1:]...)
		}
	}
}

// Not done
func (ac *ActionsContainer) AddContainerAction(container ActionsContainerInterface) {
	titleGroupMap := make(map[string]ActionsContainerInterface)
	for _, groupInterface := range ac.ActionGroups {
		group := groupInterface.(*ArgumentGroup)
		if _, found := titleGroupMap[group.Title]; found {
			panic(fmt.Sprintf("cannot merge actions - two groups are named %v", group.Title))
		}
		titleGroupMap[group.Title] = group
	}

	//  map each action to its group
	groupMap := make(map[ActionInterface]ActionsContainerInterface)
	for _, groupInterface := range container.Struct().ActionGroups {
		group := groupInterface.(*ArgumentGroup)
		// if a group with the title exists, use that, otherwise
		// create a new group matching the container's group
		if _, found := titleGroupMap[group.Title]; found {
			ac.AddArgumentGroup(
				NewArgumentGroup(
					ac,
					group.Title,
					group.Description,
					"",
					group.ConflictHandler,
					nil,
				),
			)
		}

		for _, action := range groupInterface.(*ArgumentGroup).GroupActions {
			groupMap[action] = titleGroupMap[group.Title]
		}
	}

	// add container's mutually exclusive groups
	// NOTE: if add_mutually_exclusive_group ever gains title= and
	// description= then this code will need to be expanded as above
	var cont ActionsContainerInterface
	for _, groupInterface := range container.Struct().MutuallyExclusiveGroups {
		group := groupInterface.(*MutuallyExclusiveGroup)
		if group.ActionsContainer == container {
			cont = ac
		} else {
			cont = titleGroupMap[group.ActionsContainer.Title]
		}
		mutexGroup := cont.AddMutuallyExclusiveGroup(
			&MutuallyExclusiveGroup{
				ActionsContainer: &ActionsContainer{
					Required: group.Required,
				},
			},
		)

		// map the actions to their new mutex group
		for _, action := range group.GroupActions {
			if _, found := groupMap[action]; found {
				groupMap[action] = mutexGroup
			}
		}
	}

	// add all actions to this container or their group
	for _, action := range container.Struct().Actions {
		if group, found := groupMap[action]; found {
			group.AddAction(action)
		} else {
			ac.AddAction(action)
		}
	}
}

// _get_positional_kwargs
func (ac *ActionsContainer) GetPositionalArgument(argument *Argument) *Argument {
	// make sure required is not specified
	if argument.Required {
		panic("'Required' is an invalid argument for positionals")
	}

	// mark positional arguments as required if at least one is
	// always required
	nargs := argument.Nargs
	switch t := nargs.(type) {
	case int:
		if t == 0 { // `t` is the asserted int value
			panic("nargs for positionals must be != 0")
		}
	case string:
		if t != OPTIONAL && t != ZERO_OR_MORE && t != REMAINDER && t != SUPPRESS {
			argument.Required = true
		}
	}

	// return the keyword arguments with no option strings
	argument.Dest = argument.OptionStrings[0]
	argument.OptionStrings = []string{}
	return argument
}

// _get_optional_kwargs
func (ac *ActionsContainer) GetOptionalArgument(argument *Argument) *Argument {
	optionStrings := []string{}
	longOptionStrings := []string{}
	for _, optionString := range argument.OptionStrings {
		if !isArgInChars(optionString[:1], ac.PrefixChars) {
			panic(fmt.Sprintf("invalid option string %v: must start with a character %s", optionString, ac.PrefixChars))
		}

		// strings starting with two prefix characters are long options
		optionStrings = append(optionStrings, optionString)
		if len(optionString) > 1 && isArgInChars(optionString[1:2], ac.PrefixChars) {
			longOptionStrings = append(longOptionStrings, optionString)
		}
	}

	// infer destination, '--foo-bar' -> 'foo_bar' and '-x' -> 'x'
	dest := argument.Dest
	var destOptionString string

	if dest == "" {
		if len(longOptionStrings) > 1 {
			destOptionString = longOptionStrings[0]
		} else {
			destOptionString = optionStrings[0]
		}
		dest = strings.TrimLeft(destOptionString, ac.PrefixChars)
		if dest == "" {
			panic(fmt.Sprintf("dest= is required for options like %v", optionStrings))
		}
		dest = strings.ReplaceAll(dest, "-", "_")
	}

	// return the updated argument
	argument.Dest = dest
	argument.OptionStrings = optionStrings
	return argument
}

// FIXME: not done
func (ac *ActionsContainer) GetHandler() func(ActionInterface, []ConflictingOption) {

	switch ac.ConflictHandler {
	// case "ignore":
	// 	return func() {
	// 		fmt.Println("Ignoring conflicts")
	// 	}
	// case "resolve":
	// 	return func() {
	// 		fmt.Println("Resolving conflicts")
	// 	}
	default:
		return func(ai ActionInterface, a []ConflictingOption) {} // temp
		// panic(fmt.Sprintf("invalid conflict resolution value: %v", ac.ConflictHandler))
	}
}

func (ac *ActionsContainer) CheckConflict(action ActionInterface) {

	// find all options that conflict with this option
	var conflOptionals []ConflictingOption
	for _, optionString := range action.Struct().OptionStrings {
		if _, found := ac.OptionStringActions[optionString]; found {
			conflOptional := ac.OptionStringActions[optionString]
			conflOptionals = append(conflOptionals, ConflictingOption{
				OptionString:   optionString,
				ConflictAction: conflOptional,
			})
		}
	}

	// resolve any conflicts
	if len(conflOptionals) > 0 {
		conflictHandler := ac.GetHandler()
		conflictHandler(action, conflOptionals)
	}
}

func (ac *ActionsContainer) HandleConflictError(action ActionInterface, conflictingActions []ConflictingOption) {
	message := "conflicting option strings: %s"
	conflictStrings := []string{}
	for _, conflict := range conflictingActions {
		optionString := conflict.OptionString
		action = conflict.ConflictAction
		conflictStrings = append(conflictStrings, optionString)
	}
	conflictString := fmt.Sprintf(message, strings.Join(conflictStrings, ", "))
	panic(fmt.Sprintf("%v: %s", action, conflictString))
}

func (ac *ActionsContainer) HandleConflictResolve(action ActionInterface, conflictingActions []ConflictingOption) {
	// remove all conflicting options

	for _, item := range conflictingActions {
		optionString := item.OptionString
		action := item.ConflictAction

		slice := action.Struct().OptionStrings

		for i, v := range slice {
			if v == optionString {
				// Remove the item by slicing the array
				action.Struct().OptionStrings = append(slice[:i], slice[i+1:]...)
				delete(ac.OptionStringActions, optionString)

				// if the option now has no option string, remove it from the
				// container holding it
				if len(action.Struct().OptionStrings) == 0 {
					action.Struct().Container.(*ActionsContainer).RemoveAction(action)
				}
			}
		}

	}
}

// FIXME: not done
func (ac *ActionsContainer) CheckHelp(action ActionInterface) {
	// formatter := action.Struct().GetFormatter
	// if action.Struct().Help != "" && formatter {
	// 	formatter.ExpandHelp(action)
	// }
}

// func (ac *ActionsContainer) createAction(actionName any, argument *Argument) ActionInterface {
// 	// create the action object, and add it to the parser
// 	createAction := ac.registryGet("action", actionName, actionName)

// 	callbackVal := reflect.ValueOf(createAction)
// 	if callbackVal.Kind() != reflect.Func {
// 		panic(fmt.Sprintf("unknown action: %v: %v", actionName, createAction))
// 	}

// 	// Prepare the arguments for the function call
// 	argsVals := []reflect.Value{reflect.ValueOf(argument)}

// 	// Call the function with the prepared arguments
// 	resultVals := callbackVal.Call(argsVals)

// 	// Get the first result (assumes the function returns one value)
// 	result := resultVals[0].Interface()

// 	// Перевіряємо, чи результат реалізує інтерфейс ActionInterface
// 	if action, ok := result.(ActionInterface); ok {
// 		return action
// 	}

// 	panic(fmt.Sprintf("result does not implement ActionInterface: %T", result))

// 	// // If result is already *Action, return it
// 	// if action, ok := result.(*Action); ok {
// 	// 	return action
// 	// }

// 	// // If result embeds Action (value), use reflection to find and return a pointer to it
// 	// val := reflect.ValueOf(result)
// 	// if val.Kind() == reflect.Ptr {
// 	// 	val = val.Elem() // Dereference pointer
// 	// }

// 	// // Search for embedded Action field
// 	// for i := 0; i < val.NumField(); i++ {
// 	// 	field := val.Field(i)
// 	// 	if field.Type() == reflect.TypeOf(Action{}) {
// 	// 		// Get a pointer to the embedded Action
// 	// 		return field.Addr().Interface().(*Action)
// 	// 	}
// 	// }

// 	// // If no Action is found, panic
// 	// panic(fmt.Sprintf("result does not embed Action: %T", result))

// }
