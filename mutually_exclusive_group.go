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

func (ag *MutuallyExclusiveGroup) AddArgumentGroup(argumentGroup ActionsContainerInterface) ActionsContainerInterface {
	panic("argument group can not be nested")
}

func (a *MutuallyExclusiveGroup) AddMutuallyExclusiveGroup(mutuallyExclusiveGroup ActionsContainerInterface) ActionsContainerInterface {
	panic("mutually exclusive groups cannot be nested")
}
