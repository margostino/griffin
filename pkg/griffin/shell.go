package griffin

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/margostino/griffin/pkg/action"
	"github.com/margostino/griffin/pkg/command"
	"github.com/margostino/griffin/pkg/config"
	"github.com/margostino/griffin/pkg/context"
	"log"
	"strconv"
	"strings"
)

type Shell struct {
	Suggestions        []prompt.Suggest
	CommandMap         *command.CommandMap
	ActionMap          map[string]func()
	ActionOneStringMap map[string]func([]string)
}

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
func (s *Shell) GetOptions() []prompt.Suggest {
	return s.Suggestions
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

func getMetadata(commandMap map[string]*command.Command) []prompt.Suggest {
	var suggestions = make([]prompt.Suggest, 0)
	for key, value := range commandMap {
		var commandText string
		if value.Args > 0 {
			commandText = key + " x" + strconv.Itoa(value.Args)
		} else {
			commandText = key
		}
		suggestion := prompt.Suggest{
			Text:        commandText,
			Description: value.Description,
		}
		suggestions = append(suggestions, suggestion)
	}
	return suggestions
}

func New() *Shell {
	return &Shell{Suggestions: make([]prompt.Suggest, 0), CommandMap: nil, ActionMap: nil, ActionOneStringMap: nil}
}

func Completer(suggestions []prompt.Suggest) func(d prompt.Document) []prompt.Suggest {
	return func(d prompt.Document) []prompt.Suggest {
		return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
	}
}

func commandBind(commandsList []config.CommandConfiguration, shell *Shell) *command.CommandMap {
	commands := make(map[string]*command.Command)

	for _, value := range commandsList {
		action := getAction(&value, shell)
		if action != nil && (isValidAction(&value, action.Function) || isValidInputAction(&value, action.InputFunction)) {
			commands[value.Id] = &command.Command{
				Id:          value.Id,
				Args:        value.Args,
				Action:      getAction(&value, shell),
				Pattern:     value.Pattern,
				Description: value.Description,
			}
		} else {
			log.Printf("Command %s with Args %d, Pattern %s and action %s is not valid\n", value.Id, value.Args, value.Pattern, value.Action)
		}

	}

	return newCommandMap(commands)
}

func getAction(command *config.CommandConfiguration, shell *Shell) *action.Action {
	var commandAction *action.Action = nil
	if command.Args > 0 {
		function := shell.ActionOneStringMap[command.Action]
		commandAction = action.NewInputAction(function)
	} else {
		function := shell.ActionMap[command.Action]
		commandAction = action.NewAction(function)
	}
	return commandAction
}

func newCommandMap(commands map[string]*command.Command) *command.CommandMap {
	return &command.CommandMap{Commands: commands}
}

func isValidInputAction(command *config.CommandConfiguration, function func([]string)) bool {
	return command.Args > 0 && command.Pattern != "" && function != nil
}

func isValidAction(command *config.CommandConfiguration, function func()) bool {
	return command.Args == 0 && command.Pattern == "" && function != nil
}
