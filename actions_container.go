package argparse

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

type Registry = map[any]any
type Registries = map[string]Registry

type ActionsContainer struct {
	Description                string
	PrefixChars                string
	ArgumentDefault            any
	ConflictHandler            any
	registries                 Registries
	Actions                    []ActionInterface
	optionStringActions        map[string]any
	actionGroups               []any
	mutuallyExclusiveGroups    []any
	defaults                   map[string]any
	negativeNumberMatcher      *regexp.Regexp
	hasNegativeNumberOptionals []bool

	getFormatter any
}

func NewActionsContainer(
	description string,
	prefixChars string,
	argumentDefault any,
	conflictHandler any,
) *ActionsContainer {

	container := &ActionsContainer{
		Description:                description,
		PrefixChars:                prefixChars,
		ArgumentDefault:            argumentDefault,
		ConflictHandler:            conflictHandler,
		registries:                 make(Registries),    // set up registries
		Actions:                    []ActionInterface{}, // action storage
		optionStringActions:        make(map[string]any),
		actionGroups:               []any{}, // groups
		mutuallyExclusiveGroups:    []any{},
		defaults:                   make(map[string]any),            // defaults storage
		negativeNumberMatcher:      regexp.MustCompile(`-\.?(\d+)`), // determines whether an "option" looks like a negative number
		hasNegativeNumberOptionals: []bool{},                        // # whether or not there are any optionals that look like negative numbers -- uses a list so it can be shared and edited
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
	registry, exists := ac.registries[registryName]
	if !exists {
		registry = make(Registry)
		ac.registries[registryName] = registry
	}

	// Add the value-object pair to the registry
	registry[value] = object
}

// _RegistryGet method to retrieve a value from a registry with a default
func (ac *ActionsContainer) registryGet(registryName string, value any, defaultVal any) any {
	// Retrieve the registry by name
	registry, exists := ac.registries[registryName]
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
func (ac *ActionsContainer) SetDefaults(kwargs map[string]any) {
	// Update the _defaults map
	for key, value := range kwargs {
		ac.defaults[key] = value
	}

	// Update the default value of actions that match the keys in kwargs
	for _, actionInterface := range ac.Actions {
		action := actionInterface.Self()
		if defaultValue, exists := kwargs[action.Dest]; exists {
			action.Default = defaultValue
		}
	}
}

func (ac *ActionsContainer) GetDefault(dest string) any {
	// Iterate over actions to find a matching dest with a non-nil default value
	for _, actionInterface := range ac.Actions {
		action := actionInterface.Self()
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
	Action        string
}

func isArgInChars(value string, chars string) bool {
	for _, c := range chars {
		if string(c) == value {
			return true
		}
	}
	return false
}

func (ac *ActionsContainer) AddArgument(argument *Argument) any {
	chars := ac.PrefixChars
	args := argument.OptionStrings

	// Check if args are empty or not of a valid prefix
	if args == nil || (len(args) == 1) && !isArgInChars(args[0][:1], chars) {
		// Handle positional arguments
		if argument.Dest != "" {
			panic("Dest supplied twice for positional argument, did you mean Metavar")
		}
		argument = ac.getPositionalArgument(argument)
	} else {
		// Handle optional arguments
		argument = ac.getOptionalArgument(argument)
	}

	// if no default was supplied, use the parser-level default
	if argument.Default == nil {
		dest := argument.Dest
		if _, exist := ac.defaults[dest]; exist {
			argument.Default = ac.defaults[dest]
		} else if ac.ArgumentDefault != nil {
			argument.Default = ac.ArgumentDefault
		}
	}

	// Check if the registry value is of type NewActionFuncType and invoke it
	actionName := argument.Action

	action := ac.createAction(actionName, argument)

	// fmt.Println(a)
	// var action *Action

	// if createAction, ok := newActionFunc.(func(*Argument) *Action); ok {
	// 	action = createAction(argument)
	// 	// if action, ok = createAction(argument).(*Action); !ok {
	// 	// 	panic(fmt.Sprintf("unknown action: %v: %v", actionName, newActionFunc))
	// 	// }
	// } else {
	// 	panic(fmt.Sprintf("unknown action: %v: %v", actionName, newActionFunc))
	// }

	// raise an error if action for positional argument does not consume arguments
	if action.Self().OptionStrings == nil {
		if nargs, ok := action.Self().Nargs.(int); ok && nargs == 0 {
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
	// // raise an error if the metavar does not match the type
	// if ac.getFormatter != nil {
	// 	formatter := ac.getFormatter.(FormatterFunc)()
	// 	if err := formatter.formatArgs(action, nil); err {
	// 		panic("length of metavar tuple does not match nargs")
	// 	}
	// }
	ac.CheckHelp(action)

	return ac.AddAction(action)
}

func (ac *ActionsContainer) createAction(actionName any, argument *Argument) ActionInterface {
	// create the action object, and add it to the parser
	createAction := ac.registryGet("action", actionName, actionName)

	callbackVal := reflect.ValueOf(createAction)
	if callbackVal.Kind() != reflect.Func {
		panic(fmt.Sprintf("unknown action: %v: %v", actionName, createAction))
	}

	// Prepare the arguments for the function call
	argsVals := []reflect.Value{reflect.ValueOf(argument)}

	// Call the function with the prepared arguments
	resultVals := callbackVal.Call(argsVals)

	// Get the first result (assumes the function returns one value)
	result := resultVals[0].Interface()

	// Перевіряємо, чи результат реалізує інтерфейс ActionInterface
	if action, ok := result.(ActionInterface); ok {
		return action
	}

	panic(fmt.Sprintf("result does not implement ActionInterface: %T", result))

	// // If result is already *Action, return it
	// if action, ok := result.(*Action); ok {
	// 	return action
	// }

	// // If result embeds Action (value), use reflection to find and return a pointer to it
	// val := reflect.ValueOf(result)
	// if val.Kind() == reflect.Ptr {
	// 	val = val.Elem() // Dereference pointer
	// }

	// // Search for embedded Action field
	// for i := 0; i < val.NumField(); i++ {
	// 	field := val.Field(i)
	// 	if field.Type() == reflect.TypeOf(Action{}) {
	// 		// Get a pointer to the embedded Action
	// 		return field.Addr().Interface().(*Action)
	// 	}
	// }

	// // If no Action is found, panic
	// panic(fmt.Sprintf("result does not embed Action: %T", result))

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

func (ac *ActionsContainer) AddAction(action ActionInterface) ActionInterface {
	// resolve any conflicts
	ac.CheckConflict(action)

	// add to actions list
	ac.Actions = append(ac.Actions, action)
	action.Self().container = ac

	// index the action by any option strings it has
	for _, optionString := range action.Self().OptionStrings {
		ac.optionStringActions[optionString] = action
	}

	// set the flag if any option strings look like negative numbers
	for _, optionString := range action.Self().OptionStrings {
		if ac.negativeNumberMatcher.MatchString(optionString) {
			if len(ac.hasNegativeNumberOptionals) == 0 {
				ac.hasNegativeNumberOptionals = append(ac.hasNegativeNumberOptionals, true)
			}
		}
	}
	// return the created action
	return action
}

func (ac *ActionsContainer) RemoveAction(action *Action) {

}

func (ac *ActionsContainer) AddContainerAction(container *ActionsContainer) {

}

// _get_positional_kwargs
func (ac *ActionsContainer) getPositionalArgument(argument *Argument) *Argument {
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
	argument.OptionStrings = []string{}
	return argument
}

// _get_optional_kwargs
func (ac *ActionsContainer) getOptionalArgument(argument *Argument) *Argument {
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

func (ac *ActionsContainer) CheckConflict(action ActionInterface) {

}

func (ac *ActionsContainer) HandleConflictError(action *Action, conflictingActions []*Action) {

}

func (ac *ActionsContainer) HandleConflictResolve(action *Action, conflictingActions []*Action) {

}

func (ac *ActionsContainer) CheckHelp(action ActionInterface) {

}
