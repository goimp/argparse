package argparse

import (
	"argparse/namespace"
)

// CountAction represents an action that counts the occurrences of a command-line option.
type CountAction struct {
	Action // Embedding Action to reuse functionality
}

// NewCountAction creates a new CountAction.
func NewCountAction(
	optionStrings []string,
	dest string,
	defaultVal any,
	required bool,
	help string,
	deprecated bool,
) (*CountAction, error) {
	return &CountAction{
		Action: Action{
			OptionStrings: optionStrings,
			Dest:          dest,
			Nargs:         0,
			Default:       defaultVal,
			Required:      required,
			Help:          help,
			Deprecated:    deprecated,
		},
	}, nil
}

func (a *CountAction) Call(parser any, namespace *namespace.Namespace, values any, optionString string) {
	count, found := namespace.Get(a.Dest)
	if !found || count == nil {
		count = 0
	}
	namespace.Set(a.Dest, count.(int)+1)
}

// func (a *CountAction) Call(parser any, namespace *namespace.Namespace, values any, optionString string) {
// 	var count int
// 	if c, found := namespace.Get(a.Dest); found {
// 		if castCount, ok := c.(int); ok {
// 			count = castCount
// 		}
// 	}
// 	namespace.Set(a.Dest, count+1)
// }
