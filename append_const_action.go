package argparse

// AppendConstAction represents an action that appends a constant value to a slice.
type AppendConstAction struct {
	*Action // Embedding Action to reuse functionality
}

// NewAppendConstAction creates a new AppendConstAction.
func NewAppendConstAction(argument *Argument) ActionInterface {
	return &AppendConstAction{
		Action: &Action{
			OptionStrings: argument.OptionStrings,
			Dest:          argument.Dest,
			Nargs:         0,
			Const:         argument.Const,
			Default:       argument.Default,
			Required:      argument.Required,
			Help:          argument.Help,
			MetaVar:       argument.MetaVar,
			Deprecated:    argument.Deprecated,
		},
	}
}

// func (a *AppendConstAction) Call(parser *ArgumentParser, namespace *Namespace, values any, optionString string) error {
// 	items, found := namespace.Get(a.Dest)
// 	if !found {
// 		items = []any{}
// 	}
// 	items = copy_items.CopyItems(items)
// 	items = append(items.([]any), a.Const)
// 	namespace.Set(a.Dest, items)
// 	return nil
// }

func (a *AppendConstAction) Call(parser *ArgumentParser, namespace *Namespace, values any, optionString string) error {
	items, found := namespace.Get(a.Dest)
	if !found {
		items = []any{}
	}
	items = CopyItems(items)
	items = append(items.([]any), a.Const)
	namespace.Set(a.Dest, items)
	return nil
}
