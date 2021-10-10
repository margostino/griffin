package shell

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/margostino/griffin/pkg/command"
	"github.com/margostino/griffin/pkg/config"
	"github.com/margostino/griffin/pkg/context"
	"strconv"
	"strings"
)

type Shell struct {
	Suggestions []prompt.Suggest
	CommandMap  *command.CommandMap
}

var PowerShell *Shell

func (s *Shell) Prompt() string {
	prefix := fmt.Sprintf("%s> ", "griffin")
	return prompt.Input(strings.ToLower(prefix), Completer(s.Suggestions))
}

func (s *Shell) Input() string {
	commandLine := s.Prompt()
	return commandLine
}

func (s *Shell) Start() {
	var plan string
	for {
		plan = s.Input()
		command := s.CommandMap.Lookup(plan)
		if command != nil {
			context.Prepare(plan).With(command).Execute()
		} else {
			fmt.Printf("command plan %q is not valid\n", plan)
		}
	}
}

func NewShell(configFile string) *Shell {
	var suggestions = make([]prompt.Suggest, 0)
	commandConfig := config.LoadCommands(configFile)
	commands := commandConfig.CommandList
	commandMap := command.GetCommandMap(commands)

	for _, command := range commands {
		var commandText string
		if command.Args > 0 {
			commandText = command.Id + " x" + strconv.Itoa(command.Args)
		} else {
			commandText = command.Id
		}
		suggestion := prompt.Suggest{
			Text:        commandText,
			Description: command.Description,
		}
		suggestions = append(suggestions, suggestion)
	}
	PowerShell = &Shell{Suggestions: suggestions, CommandMap: commandMap}
	return PowerShell
}

func Completer(suggestions []prompt.Suggest) func(d prompt.Document) []prompt.Suggest {
	return func(d prompt.Document) []prompt.Suggest {
		return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
	}
}

func (s *Shell) GetOptions() []prompt.Suggest {
	return s.Suggestions
}
