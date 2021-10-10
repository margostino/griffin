package command

import (
	"github.com/margostino/griffin/pkg/action"
	"github.com/margostino/griffin/pkg/config"
)

func newCommandMap(commands map[string]*Command) *CommandMap {
	return &CommandMap{commands}
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

func GetAction(command *config.CommandConfiguration) *action.Action {
	var commandAction *action.Action = nil
	if command.Args > 0 {
		function := action.InputActionStorage[command.Action]
		commandAction = action.NewInputAction(function)
	} else {
		function := action.ActionStorage[command.Action]
		commandAction = action.NewAction(function)
	}
	return commandAction
}

func GetCommandMap(commandsList []config.CommandConfiguration) *CommandMap {
	commands := make(map[string]*Command)

	for _, command := range commandsList {
		commands[command.Id] = &Command{
			Id:      command.Id,
			Args:    command.Args,
			Action:  GetAction(&command),
			Pattern: command.Pattern,
		}
	}

	return newCommandMap(commands)
}
