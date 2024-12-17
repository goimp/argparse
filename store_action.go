package argparse

import (
	"fmt"
)

type StoreAction struct {
	*Action
}

func NewStoreAction(argument *Argument) *StoreAction {

	switch v := argument.Nargs.(type) {
	case int:
		if v == 0 {
			panic(fmt.Sprintf("nargs for store actions must be != 0; if you have nothing to store, actions such as store true or store const may be more appropriate"))
		}
	case string:
		if argument.Const != nil && v != OPTIONAL {
			panic(fmt.Sprintf("nargs must %s to supply const", OPTIONAL))
		}
	case nil:
		break
	default:
		panic(fmt.Sprintf("nargs must be an integer or a string literal (e.g., %s)", OPTIONAL))
	}

	return &StoreAction{
		Action: &Action{
			OptionStrings: argument.OptionStrings,
			Dest:          argument.Dest,
			Nargs:         argument.Nargs,
			Const:         argument.Const,
			Default:       argument.Default,
			Type:          argument.Type,
			Choices:       argument.Choices,
			Required:      argument.Required,
			Help:          argument.Help,
			Metavar:       argument.Metavar,
			Deprecated:    argument.Deprecated,
		},
	}
}

// Call assigns the provided values to the destination field in the namespace.
func (a *StoreAction) Call(parser *ArgumentParser, namespace *Namespace, values any, optionString string) error {
	namespace.Set(a.Dest, values)
	return nil
}
