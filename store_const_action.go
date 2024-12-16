package argparse

import "argparse/namespace"

// StoreConstAction represents an action that stores a constant value in the namespace.
type StoreConstAction struct {
	Action // Embed Action to inherit its behavior
}

// // NewStoreConstAction creates a new StoreConstAction with the provided parameters.
func NewStoreConstAction(optionStrings []string, dest string, constVal any, defaultVal any,
	required bool, help, metavar string, deprecated bool) (*StoreConstAction, error) {
	// Validate nargs

	// Create the base Action
	action := Action{
		OptionStrings: optionStrings,
		Dest:          dest,
		Nargs:         0,
		Const:         constVal,
		Default:       defaultVal,
		Required:      required,
		Help:          help,
		Metavar:       metavar,
		Deprecated:    deprecated,
	}

	return &StoreConstAction{Action: action}, nil
}

// Call assigns the provided values to the destination field in the namespace.
func (a *StoreConstAction) Call(parser *ArgumentParser, namespace *namespace.Namespace, values any, optionString string) {
	namespace.Set(a.Dest, a.Const)
}
