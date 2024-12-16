package argparse

type ArgumentGroup struct {
	ActionsContainer
}

func NewArgumentGroup(
	description string,
	prefixChars any,
	argumentDefault any,
	conflictHandler any,
) (*ArgumentGroup, error) {

	actions_container, error := NewActionsContainer(
		description,
		prefixChars,
		argumentDefault,
		conflictHandler,
	)

	if error != nil {
		return nil, error
	}

	return &ArgumentGroup{
		ActionsContainer: *actions_container,
	}, nil
}

func (a *ArgumentGroup) AddAction(action Action) *Action {
	return &Action{}
}

func (a *ArgumentGroup) RemoveAction(action Action) {

}

func (a *ArgumentGroup) AddArgumentGroup(args []any, kwargs map[string]any) {
	panic("argument group can not be nested")
}
