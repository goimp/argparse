package action

// StoreConstAction represents an action that stores a constant value in the namespace.
type StoreConstAction struct {
	OptionStrings []string
	Dest          string
	Const         interface{}
	Default       interface{}
	Required      bool
	Help          string
	Metavar       interface{}
	Deprecated    bool
}

// NewStoreConstAction creates a new StoreConstAction with the provided parameters.
func NewStoreConstAction(optionStrings []string, dest string, constVal interface{}, defaultVal interface{}, help string) *StoreConstAction {
	return &StoreConstAction{
		OptionStrings: optionStrings,
		Dest:          dest,
		Const:         constVal,
		Default:       defaultVal,
		Help:          help,
	}
}
