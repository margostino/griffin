package griffin

import (
	"github.com/margostino/griffin/pkg/config"
	"testing"
)

func TestLoadConfiguration(t *testing.T) {
	powershell := New().LoadConfiguration("../../example/config/commands.yml")
	if powershell.CommandMap.Commands == nil {
		t.Fatalf(`it should load`)
	}
}

func TestFailedLoadConfiguration(t *testing.T) {
	powershell := New().LoadConfiguration("invalid")
	if len(powershell.CommandMap.Commands) > 0 {
		t.Fatalf(`it should not load`)
	}
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
		SetActionMap(ActionMap).
		SetConfiguration(&commands)

	if len(powershell.CommandMap.Commands) != 1 {
		t.Fatalf(`it should load one command`)
	}

	if powershell.CommandMap.Commands["run test"] == nil {
		t.Fatalf(`it should load dummy command`)
	}

	if powershell.CommandMap.Commands["run test"].Args != 0 {
		t.Fatalf(`it should load dummy command with args 0`)
	}

	if powershell.CommandMap.Commands["run test"].Pattern != "" {
		t.Fatalf(`it should load dummy command with no pattern`)
	}

	if powershell.CommandMap.Commands["run test"].Description != "testing commands" {
		t.Fatalf(`it should load dummy command with description`)
	}

	if powershell.CommandMap.Commands["run test"].Action.Function == nil {
		t.Fatalf(`it should load dummy command with action`)
	}
}
