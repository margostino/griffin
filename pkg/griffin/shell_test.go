package griffin

import (
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
	dummyCommand := CommandConfiguration{
		Id:          "run test",
		Description: "testing commands",
		Args:        0,
		Action:      "ExecuteDummyAction",
	}
	commands := CommandsConfiguration{
		CommandList: []CommandConfiguration{dummyCommand},
	}
	powershell := New().
		SetSimpleActions(ActionMap).
		SetConfiguration(&commands)

	assertShouldLoadCommand(powershell, "run test", 0, "", "testing commands", t)
	assertShouldLoadCommandAction(powershell, "run test", t)
	assertShouldNotLoadCommandInputAction(powershell, "run test", t)
}

func TestSetActions(t *testing.T) {
	dummyCommand := CommandConfiguration{
		Id:          "run test",
		Description: "testing commands",
		Args:        0,
		Action:      "ExecuteDummyAction",
	}
	commands := CommandsConfiguration{
		CommandList: []CommandConfiguration{dummyCommand},
	}
	powershell := New().
		SetSimpleActions(ActionMap).
		SetConfiguration(&commands)

	assertShouldLoadCommand(powershell, "run test", 0, "", "testing commands", t)

	if powershell.CommandMap.Commands["run test"].Action.Function == nil {
		t.Fatalf(`it should load dummy command with action`)
	}
}

func TestSetInputActions(t *testing.T) {
	dummyCommand := CommandConfiguration{
		Id:          "run test",
		Description: "testing commands",
		Args:        1,
		Action:      "ExecuteDummyInputAction",
		Pattern:     "^run test [a-z-A-Z]+$",
	}
	commands := CommandsConfiguration{
		CommandList: []CommandConfiguration{dummyCommand},
	}
	powershell := New().
		SetMultiParamsActions(ActionOneString).
		SetConfiguration(&commands)

	assertShouldLoadCommand(powershell, "run test", 1, "^run test [a-z-A-Z]+$", "testing commands", t)
	assertShouldLoadCommandInputAction(powershell, "run test", t)
	assertShouldNotLoadCommandAction(powershell, "run test", t)
}

func TestEmptyPatternWithActionHasArgs(t *testing.T) {
	dummyCommand := CommandConfiguration{
		Id:          "run test",
		Description: "testing commands",
		Args:        1,
		Action:      "ExecuteDummyInputAction",
		Pattern:     "",
	}
	commands := CommandsConfiguration{
		CommandList: []CommandConfiguration{dummyCommand},
	}
	powershell := New().
		SetMultiParamsActions(ActionOneString).
		SetConfiguration(&commands)

	assertShouldNotLoadCommandsWithInvalidConfig(powershell, t)
}

func TestNonEmptyPatternWithoutArgs(t *testing.T) {
	dummyCommand := CommandConfiguration{
		Id:          "run test",
		Description: "testing commands",
		Args:        0,
		Action:      "ExecuteDummyInputAction",
		Pattern:     "^run test [a-z-A-Z]+$",
	}
	commands := CommandsConfiguration{
		CommandList: []CommandConfiguration{dummyCommand},
	}
	powershell := New().
		SetMultiParamsActions(ActionOneString).
		SetConfiguration(&commands)

	assertShouldNotLoadCommandsWithInvalidConfig(powershell, t)
}

func TestActionIsNotFound(t *testing.T) {
	dummyCommand := CommandConfiguration{
		Id:          "run test",
		Description: "testing commands",
		Args:        1,
		Action:      "InvalidFunction",
		Pattern:     "some regex",
	}
	commands := CommandsConfiguration{
		CommandList: []CommandConfiguration{dummyCommand},
	}
	powershell := New().
		SetMultiParamsActions(ActionOneString).
		SetConfiguration(&commands)

	assertShouldNotLoadCommandsWithInvalidConfig(powershell, t)
}

func TestSuggestion(t *testing.T) {
	dummyCommand := CommandConfiguration{
		Id:          "run test",
		Description: "testing commands",
		Args:        0,
		Action:      "ExecuteDummyAction",
	}
	commands := CommandsConfiguration{
		CommandList: []CommandConfiguration{dummyCommand},
	}
	powershell := New().
		SetSimpleActions(ActionMap).
		SetConfiguration(&commands)

	assertShouldLoadSuggestion(powershell, "run test", "testing commands", t)
}

func TestPromptName(t *testing.T) {
	powershell := New().SetPrompt("testing")
	assertPromptName(powershell, "testing", t)
}
