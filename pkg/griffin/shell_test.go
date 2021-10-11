package griffin

import (
	"github.com/margostino/griffin/pkg/config"
	"testing"
)

func TestLoadConfiguration(t *testing.T) {
	powershell := New().LoadConfiguration("../../example/config/commands.yml")
	assertShouldLoadCommands(powershell, t)
}

func TestFailedLoadConfiguration(t *testing.T) {
	powershell := New().LoadConfiguration("invalid")
	assertShouldNotLoadCommandsWithInvalidConfig(powershell, t)
	assertShouldNotLoadSuggestions(powershell, t)
	assertShouldNotLoadActions(powershell, t)
	assertShouldNotLoadInputActions(powershell, t)
}

func TestFailedSetConfiguration(t *testing.T) {
	powershell := New().SetConfiguration(nil)
	assertShouldNotLoadCommands(powershell, t)
	assertShouldNotLoadSuggestions(powershell, t)
	assertShouldNotLoadActions(powershell, t)
	assertShouldNotLoadInputActions(powershell, t)
}

func TestSetConfiguration(t *testing.T) {
	dummyCommand := config.CommandConfiguration{
		Id:          "run test",
		Description: "testing commands",
		Args:        0,
		Action:      "ExecuteDummyAction",
	}
	commands := config.CommandsConfiguration{
		CommandList: []config.CommandConfiguration{dummyCommand},
	}
	powershell := New().
		SetActions(ActionMap).
		SetConfiguration(&commands)

	assertShouldLoadCommand(powershell, "run test", 0, "", "testing commands", t)
	assertShouldLoadCommandAction(powershell, "run test", t)
	assertShouldNotLoadCommandInputAction(powershell, "run test", t)
}

func TestSetActions(t *testing.T) {
	dummyCommand := config.CommandConfiguration{
		Id:          "run test",
		Description: "testing commands",
		Args:        0,
		Action:      "ExecuteDummyAction",
	}
	commands := config.CommandsConfiguration{
		CommandList: []config.CommandConfiguration{dummyCommand},
	}
	powershell := New().
		SetActions(ActionMap).
		SetConfiguration(&commands)

	assertShouldLoadCommand(powershell, "run test", 0, "", "testing commands", t)

	if powershell.CommandMap.Commands["run test"].Action.Function == nil {
		t.Fatalf(`it should load dummy command with action`)
	}
}

func TestSetInputActions(t *testing.T) {
	dummyCommand := config.CommandConfiguration{
		Id:          "run test",
		Description: "testing commands",
		Args:        1,
		Action:      "ExecuteDummyInputAction",
		Pattern:     "^run test [a-z-A-Z]+$",
	}
	commands := config.CommandsConfiguration{
		CommandList: []config.CommandConfiguration{dummyCommand},
	}
	powershell := New().
		SetActionsStrings(ActionOneString).
		SetConfiguration(&commands)

	assertShouldLoadCommand(powershell, "run test", 1, "^run test [a-z-A-Z]+$", "testing commands", t)
	assertShouldLoadCommandInputAction(powershell, "run test", t)
	assertShouldNotLoadCommandAction(powershell, "run test", t)
}
