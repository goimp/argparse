package argparse

// StoreFalseAction represents an action that stores a constant false value in the namespace.

type StoreFalseAction struct {
	*StoreConstAction // Embed Action to inherit its behavior
}

func NewStoreFalseAction(argument *Argument) *StoreFalseAction {
	argument.Const = false
	storeConstAction := NewStoreConstAction(argument)
	return &StoreFalseAction{
		StoreConstAction: storeConstAction,
	}
}

// Make sure StoreTrueAction implements ActionInterface
func (a *StoreFalseAction) Self() *Action {
	return a.StoreConstAction.Self() // Call Self() from StoreConstAction
}

func (a *StoreFalseAction) GetKwargs() map[string]any {
	return a.StoreConstAction.GetKwargs()
}

func (a *StoreFalseAction) FormatUsage() string {
	return a.StoreConstAction.FormatUsage()
}

func (a *StoreFalseAction) Call(parser *ArgumentParser, namespace *Namespace, values any, optionString string) error {
	a.StoreConstAction.Call(parser, namespace, values, optionString)
	return nil
}
