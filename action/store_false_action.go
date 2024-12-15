package action

// StoreFalseAction represents an action that stores a constant false value in the namespace.
type StoreFalseAction struct {
	StoreConstAction // Embedding StoreConstAction to reuse functionality
}

func NewStoreFalseAction(optionStrings []string, dest string, defaultVal bool, required bool, help string, deprecated bool) (*StoreFalseAction, error) {
	if storeConstAction, err := NewStoreConstAction(
		optionStrings,
		dest,
		false,
		defaultVal,
		required,
		help,
		"",
		deprecated,
	); err != nil {
		return nil, err
	} else {
		return &StoreFalseAction{
			StoreConstAction: *storeConstAction,
		}, nil
	}
}
