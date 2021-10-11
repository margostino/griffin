package griffin

import "regexp"

type Command struct {
	Id      string
	Args    int
	Pattern string
	*Action
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

func (c Command) Execute() {
	if c.Action != nil {
		c.Action.Function()
	}
}

func (c Command) ExecuteWith(args []string) {
	if c.Action != nil {
		c.Action.InputFunction(args)
	}
}
