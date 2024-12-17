package argparse

type MutuallyExclusiveGroup struct {
	*ActionsContainer
	GroupActions []ActionInterface
}

func NewMutuallyExclusiveGroup(
	description string,
	prefixChars string,
	argumentDefault any,
	conflictHandler any,
) ActionsContainerInterface {

	actionsContainer := NewActionsContainer(
		description,
		prefixChars,
		argumentDefault,
		conflictHandler,
	)

	return &MutuallyExclusiveGroup{
		ActionsContainer: actionsContainer.(*ActionsContainer),
	}
}

func (a *MutuallyExclusiveGroup) AddAction(action ActionInterface) ActionInterface {
	return &Action{}
}

func (a *MutuallyExclusiveGroup) RemoveAction(action ActionInterface) {

}

func (a *MutuallyExclusiveGroup) AddMutuallyExclusiveGroup(mutuallyExclusiveGroup *MutuallyExclusiveGroup) ActionsContainerInterface {
	panic("mutually exclusive groups cannot be nested")
}
