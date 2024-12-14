package action

// StoreTrueAction represents an action that stores a constant true value in the namespace.
type StoreTrueAction struct {
	StoreConstAction // Embedding StoreConstAction to reuse functionality
}

// NewStoreTrueAction creates a new StoreTrueAction that sets the constant value to true.
func NewStoreTrueAction(optionStrings []string, dest string, defaultVal bool, help string) *StoreTrueAction {
	return &StoreTrueAction{
		StoreConstAction: *NewStoreConstAction(optionStrings, dest, true, defaultVal, help),
	}
}
