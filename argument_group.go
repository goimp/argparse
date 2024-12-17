package argparse

import "fmt"

type ArgumentGroup struct {
	*ActionsContainer
	Title           string
	Description     string
	ConflictHandler any
	GroupActions    []ActionInterface
}

func NewArgumentGroup(
	container *ActionsContainer,
	title string,
	description string,
	prefixChars string,
	conflictHandler any,
	groupActions []ActionInterface,
) ActionsContainerInterface {

	if prefixChars != "" {
		fmt.Println("The use of the undocumented 'prefixChars' parameter in ArgumentParser.AddArgumentGroup() is deprecated.")
	}

	// kwargs["conflictHandler"] = container.ConflictHandler
	// kwargs["prefixChars"] = container.PrefixChars
	// kwargs["argumentDefault"] = container.ArgumentDefault

	action := NewActionsContainer(
		description,
		container.PrefixChars,
		container.ArgumentDefault,
		container.ConflictHandler,
	)

	group := &ArgumentGroup{
		ActionsContainer: action.(*ActionsContainer),
		Title:            title,
		GroupActions:     []ActionInterface{},
	}

	group.Registries = container.Registries
	group.Actions = container.Actions
	group.OptionStringActions = container.OptionStringActions
	group.Defaults = container.Defaults
	group.HasNegativeNumberOptionals = container.HasNegativeNumberOptionals
	group.MutuallyExclusiveGroups = container.MutuallyExclusiveGroups

	return group
}

func (ag *ArgumentGroup) AddAction(action ActionInterface) ActionInterface {
	act := ag.ActionsContainer.AddAction(action)
	ag.GroupActions = append(ag.GroupActions, act)
	return act
}

func (ac *ArgumentGroup) RemoveAction(action ActionInterface) {
	for i, v := range ac.Actions {
		if v == action {
			// Remove the item by slicing the array
			ac.Actions = append(ac.Actions[:i], ac.Actions[i+1:]...)
		}
	}
}

func (ag *ArgumentGroup) AddArgumentGroup(argumentGroup ActionsContainerInterface) ActionsContainerInterface {
	panic("argument group can not be nested")
}

func (a *ArgumentGroup) AddMutuallyExclusiveGroup(mutuallyExclusiveGroup ActionsContainerInterface) ActionsContainerInterface {
	panic("mutually exclusive groups cannot be nested")
}
