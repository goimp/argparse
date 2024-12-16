package argparse

// StoreTrueAction represents an action that stores a constant true value in the namespace.
type StoreTrueAction struct {
	StoreConstAction // Embedding StoreConstAction to reuse functionality
}

func NewStoreTrueAction(optionStrings []string, dest string, defaultVal bool, required bool, help string, deprecated bool) (*StoreTrueAction, error) {
	if storeConstAction, err := NewStoreConstAction(
		optionStrings,
		dest,
		true,
		defaultVal,
		required,
		help,
		"",
		deprecated,
	); err != nil {
		return nil, err
	} else {
		return &StoreTrueAction{
			StoreConstAction: *storeConstAction,
		}, nil
	}
}
