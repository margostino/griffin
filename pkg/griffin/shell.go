package griffin

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/margostino/griffin/pkg/command"
	"github.com/margostino/griffin/pkg/config"
	"github.com/margostino/griffin/pkg/context"
	"log"
	"strings"
)

type Shell struct {
	Suggestions        []prompt.Suggest
	CommandMap         *command.CommandMap
	ActionMap          map[string]func()
	ActionOneStringMap map[string]func([]string)
}

func (s *Shell) Input() string {
	commandLine := s.prompt()
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

func (s *Shell) LoadConfiguration(configFile string) *Shell {
	configuration := config.LoadCommands(configFile)
	return s.ConfigMap(configuration)
}

func (s *Shell) SetConfiguration(configuration *config.CommandsConfiguration) *Shell {

	if configuration == nil {
		log.Println("Configuration can not be nil")
	} else {
		commands := configuration.CommandList
		commandMap := commandBind(commands, s)
		suggestions := getMetadata(commandMap.Commands)
		s.Suggestions = suggestions
		s.CommandMap = commandMap
	}
	return s
}

func (s *Shell) SetActions(actions map[string]func()) *Shell {
	s.ActionMap = actions
	return s
}

func (s *Shell) SetActionsStrings(actions map[string]func([]string)) *Shell {
	s.ActionOneStringMap = actions
	return s
}

func (s *Shell) ConfigMap(configuration *config.CommandsConfiguration) *Shell {
	commands := configuration.CommandList
	commandMap := commandBind(commands, s)
	suggestions := getMetadata(commandMap.Commands)
	s.Suggestions = suggestions
	s.CommandMap = commandMap
	return s
}

func (s *Shell) Help() {
	for _, option := range s.Suggestions {
		fmt.Printf("[ %s ] - %s\n", option.Text, option.Description)
	}
}

func (s *Shell) prompt() string {
	prefix := fmt.Sprintf("%s> ", "griffin")
	return prompt.Input(strings.ToLower(prefix), completer(s.Suggestions))
}
