package argparse

// StoreConstAction represents an action that stores a constant value in the namespace.
type StoreConstAction struct {
	*Action // Embed Action to inherit its behavior
}

// // NewStoreConstAction creates a new StoreConstAction with the provided parameters.
func NewStoreConstAction(argument *Argument) *StoreConstAction {
	// Validate nargs

	// Create the base Action
	action := &Action{
		OptionStrings: argument.OptionStrings,
		Dest:          argument.Dest,
		Nargs:         0,
		Const:         argument.Const,
		Default:       argument.Default,
		Required:      argument.Required,
		Help:          argument.Help,
		Metavar:       argument.Metavar,
		Deprecated:    argument.Deprecated,
	}

	return &StoreConstAction{Action: action}
}

// // Make sure StoreTrueAction implements ActionInterface
// func (a *StoreConstAction) Self() *Action {
// 	return a.Action.Self() // Call Self() from StoreConstAction
// }

// func (a *StoreConstAction) GetKwargs() map[string]any {
// 	return a.Action.GetKwargs()
// }

// func (a *StoreConstAction) FormatUsage() string {
// 	return a.Action.FormatUsage()
// }

// Call assigns the provided values to the destination field in the namespace.
func (a *StoreConstAction) Call(parser *ArgumentParser, namespace *Namespace, values any, optionString string) error {
	namespace.Set(a.Dest, a.Const)
	return nil
}
