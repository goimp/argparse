package argparse

// CountAction represents an action that counts the occurrences of a command-line option.
type CountAction struct {
	*Action // Embedding Action to reuse functionality
}

// NewCountAction creates a new CountAction.
func NewCountAction(argument *Argument) ActionInterface {
	return &CountAction{
		Action: &Action{
			OptionStrings: argument.OptionStrings,
			Dest:          argument.Dest,
			Nargs:         0,
			Default:       argument.Default,
			Required:      argument.Required,
			Help:          argument.Help,
			Deprecated:    argument.Deprecated,
		},
	}
}

func (a *CountAction) Call(parser *ArgumentParser, namespace *Namespace, values any, optionString string) error {
	count, found := namespace.Get(a.Dest)
	if !found || count == nil {
		count = 0
	}
	namespace.Set(a.Dest, count.(int)+1)
	return nil
}
