package argparse

type MutuallyExclusiveGroup struct {
	ActionsContainer
}

func NewMutuallyExclusiveGroup(
	description string,
	prefixChars []string,
	argumentDefault any,
	conflictHandler any,
) (*MutuallyExclusiveGroup, error) {

	actions_container, error := NewActionsContainer(
		description,
		prefixChars,
		argumentDefault,
		conflictHandler,
	)

	if error != nil {
		return nil, error
	}

	return &MutuallyExclusiveGroup{
		ActionsContainer: *actions_container,
	}, nil
}

func (a *MutuallyExclusiveGroup) AddAction(action Action) *Action {
	return &Action{}
}

func (a *MutuallyExclusiveGroup) RemoveAction(action Action) {

}

func (a *MutuallyExclusiveGroup) AddMutuallyExclusiveGroup(args []any, kwargs map[string]any) {
	panic("mutually exclusive groups cannot be nested")
}
