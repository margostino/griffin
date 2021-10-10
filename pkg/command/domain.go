package command

import (
	"github.com/margostino/griffin/pkg/action"
	"regexp"
)

type Command struct {
	Id      string
	Args    int
	Pattern string
	*action.Action
	Description string
}

type CommandMap struct {
	Commands map[string]*Command
}

func (m *CommandMap) Lookup(plan string) *Command {
	if len(plan) == 0 {
		return nil
	}

	command, _ := m.Commands[plan]

	if command == nil {
		for _, value := range m.Commands {
			if value.Pattern != "" {
				match, _ := regexp.MatchString(value.Pattern, plan)
				if match {
					return value
				}
			}
		}
	}

	return command
}
