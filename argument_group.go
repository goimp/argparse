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
	kwargs map[string]any,
	groupActions []ActionInterface,
) ActionsContainerInterface {

	if _, exist := kwargs["prefixChars"]; exist {
		fmt.Println("The use of the undocumented 'prefix_chars' parameter in ArgumentParser.AddArgumentGroup() is deprecated.")
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

	group.registries = container.registries
	group.Actions = container.Actions
	group.optionStringActions = container.optionStringActions
	group.defaults = container.defaults
	group.hasNegativeNumberOptionals = container.hasNegativeNumberOptionals
	group.mutuallyExclusiveGroups = container.mutuallyExclusiveGroups

	return group
}

func (ag *ArgumentGroup) AddAction(action ActionInterface) ActionInterface {
	act := ag.ActionsContainer.AddAction(action)
	ag.GroupActions = append(ag.GroupActions, act)
	return act
}

func (ag *ArgumentGroup) RemoveAction(dest string) error {
	for i, actionInterface := range ag.Actions {
		action := actionInterface.Struct()
		if action.Dest == dest {
			// Remove the action by appending slices before and after the index
			ag.Actions = append(ag.Actions[:i], ag.Actions[i+1:]...)
			return nil // Action removed successfully
		}
	}
	return fmt.Errorf("action with dest %q not found", dest)
}

func (ag *ArgumentGroup) AddArgumentGroup(args []any, kwargs map[string]any) ActionsContainerInterface {
	panic("argument group can not be nested")
}
