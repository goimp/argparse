package argparse

// StoreTrueAction represents an action that stores a constant true value in the namespace.
type StoreTrueAction struct {
	*StoreConstAction // Embed Action to inherit its behavior
}

func NewStoreTrueAction(argument *Argument) *StoreTrueAction {
	argument.Const = true
	storeConstAction := NewStoreConstAction(argument)
	return &StoreTrueAction{
		StoreConstAction: storeConstAction,
	}
}

// Make sure StoreTrueAction implements ActionInterface
func (a *StoreTrueAction) Self() *Action {
	return a.StoreConstAction.Self() // Call Self() from StoreConstAction
}

func (a *StoreTrueAction) GetKwargs() map[string]any {
	return a.StoreConstAction.GetKwargs()
}

func (a *StoreTrueAction) FormatUsage() string {
	return a.StoreConstAction.FormatUsage()
}

func (a *StoreTrueAction) Call(parser *ArgumentParser, namespace *Namespace, values any, optionString string) error {
	return a.StoreConstAction.Call(parser, namespace, values, optionString)
}
