package argparse

// StoreTrueAction represents an action that stores a constant true value in the namespace.
type StoreTrueAction struct {
	*StoreConstAction // Embed Action to inherit its behavior
}

func NewStoreTrueAction(argument *Argument) ActionInterface {
	argument.Const = true
	storeConstAction := NewStoreConstAction(argument)
	return &StoreTrueAction{
		StoreConstAction: storeConstAction.(*StoreConstAction),
	}
}

// // Make sure StoreTrueAction implements ActionInterface
// func (a *StoreTrueAction) Struct() *Action {
// 	return a.StoreConstAction.Struct() // Call Struct() from StoreConstAction
// }

// func (a *StoreTrueAction) GetMap() map[string]any {
// 	return a.StoreConstAction.GetMap()
// }

// func (a *StoreTrueAction) FormatUsage() string {
// 	return a.StoreConstAction.FormatUsage()
// }

// func (a *StoreTrueAction) Call(parser *ArgumentParser, namespace *Namespace, values any, optionString string) error {
// 	return a.StoreConstAction.Call(parser, namespace, values, optionString)
// }

// func (a *StoreTrueAction) GetSubActions() []ActionInterface {
// 	return a.StoreConstAction.GetSubActions()
// }
