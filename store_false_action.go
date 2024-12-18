package argparse

// StoreFalseAction represents an action that stores a constant false value in the namespace.

type StoreFalseAction struct {
	*StoreConstAction // Embed Action to inherit its behavior
}

func NewStoreFalseAction(argument *Argument) ActionInterface {
	argument.Const = false
	storeConstAction := NewStoreConstAction(argument)
	return &StoreFalseAction{
		StoreConstAction: storeConstAction.(*StoreConstAction),
	}
}

// Make sure StoreTrueAction implements ActionInterface
func (a *StoreFalseAction) Struct() *Action {
	return a.StoreConstAction.Struct() // Call Struct() from StoreConstAction
}

func (a *StoreFalseAction) GetMap() map[string]any {
	return a.StoreConstAction.GetMap()
}

func (a *StoreFalseAction) FormatUsage() string {
	return a.StoreConstAction.FormatUsage()
}

func (a *StoreFalseAction) Call(parser *ArgumentParser, namespace *Namespace, values any, optionString string) error {
	return a.StoreConstAction.Call(parser, namespace, values, optionString)
}
