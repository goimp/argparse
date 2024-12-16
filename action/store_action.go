package action

import (
	"argparse"
	"argparse/namespace"
	"fmt"
)

// StoreAction represents an action that stores a value in the namespace.
type StoreAction struct {
	Action // Embed the Action struct
}

// NewStoreAction creates a new StoreAction object.
func NewStoreAction(optionStrings []string, dest string, nargs any, constVal any, defaultVal any,
	typ Type, choices []any, required bool, help, metavar string, deprecated bool) (*StoreAction, error) {

	switch v := nargs.(type) {
	case int:
		if v == 0 {
			return nil, fmt.Errorf("nargs for store actions must be != 0; if you have nothing to store, actions such as store true or store const may be more appropriate")
		}
	case string:
		if constVal != nil && v != argparse.OPTIONAL {
			return nil, fmt.Errorf("nargs must %s to supply const", argparse.OPTIONAL)
		}
	case nil:
		break
	default:
		return nil, fmt.Errorf("nargs must be an integer or a string literal (e.g., %s)", argparse.OPTIONAL)
	}

	action := Action{
		OptionStrings: optionStrings,
		Dest:          dest,
		Nargs:         nargs,
		Const:         constVal,
		Default:       defaultVal,
		Type:          typ,
		Choices:       choices,
		Required:      required,
		Help:          help,
		Metavar:       metavar,
		Deprecated:    deprecated,
	}

	return &StoreAction{Action: action}, nil
}

// Call assigns the provided values to the destination field in the namespace.
func (a *StoreAction) Call(parser any, namespace *namespace.Namespace, values any, optionString string) {
	namespace.Set(a.Dest, values)
}
