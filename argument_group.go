package argparse

import "fmt"

type ArgumentGroup struct {
	ActionsContainer
	title        string
	groupActions []*Action
}

func NewArgumentGroup(
	container *ActionsContainer,
	title string,
	description string,
	kwargs map[string]any,
	groupActions []*Action,
) (*ArgumentGroup, error) {

	if _, exist := kwargs["prefixChars"]; exist {
		fmt.Println("The use of the undocumented 'prefix_chars' parameter in ArgumentParser.AddArgumentGroup() is deprecated.")
	}

	// kwargs["conflictHandler"] = container.ConflictHandler
	// kwargs["prefixChars"] = container.PrefixChars
	// kwargs["argumentDefault"] = container.ArgumentDefault

	action, err := NewActionsContainer(
		description,
		container.PrefixChars,
		container.ArgumentDefault,
		container.ConflictHandler,
	)

	if err != nil {
		return nil, err
	}

	group := &ArgumentGroup{
		ActionsContainer: *action,
		title:            title,
		groupActions:     []*Action{},
	}

	group.registries = container.registries
	group.actions = container.actions
	group.optionStringActions = container.optionStringActions
	group.defaults = container.defaults
	group.hasNegativeNumberOptionals = container.hasNegativeNumberOptionals
	group.mutuallyExclusiveGroups = container.mutuallyExclusiveGroups

	return group, nil
}

func (a *ArgumentGroup) AddAction(action *Action) *Action {
	act := *a.ActionsContainer.AddAction(action)
	a.groupActions = append(a.groupActions, &act)
	return &act
}

func (a *ArgumentGroup) RemoveAction(dest string) error {
	for i, action := range a.actions {
		if action.Dest == dest {
			// Remove the action by appending slices before and after the index
			a.actions = append(a.actions[:i], a.actions[i+1:]...)
			return nil // Action removed successfully
		}
	}
	return fmt.Errorf("action with dest %q not found", dest)
}

func (a *ArgumentGroup) AddArgumentGroup(args []any, kwargs map[string]any) {
	panic("argument group can not be nested")
}
