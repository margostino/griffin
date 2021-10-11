package griffin

import (
	"fmt"
	"testing"
)

var ActionMap = map[string]func(){
	"ExecuteDummyAction": ExecuteDummyAction,
}

var ActionOneString = map[string]func([]string){
	"ExecuteDummyInputAction": ExecuteDummyInputAction,
}

func ExecuteDummyAction() {
	fmt.Println("dummy action")
}

func ExecuteDummyInputAction(args []string) {
	fmt.Printf("dummy action with input %s\n", args[0])
}

func assertShouldLoadCommands(powershell *Shell, t *testing.T) {
	if powershell.CommandMap.Commands == nil {
		t.Fatalf(`it should load command`)
	}
}

func assertShouldLoadCommand(powershell *Shell, command string, args int, pattern string, description string, t *testing.T) {
	if len(powershell.CommandMap.Commands) != 1 {
		t.Fatalf(`it should load one command`)
	}
	if powershell.CommandMap.Commands[command] == nil {
		t.Fatalf(`it should load dummy command`)
	}
	if powershell.CommandMap.Commands[command].Args != args {
		t.Fatalf(`it should load dummy command with args %d`, args)
	}
	if powershell.CommandMap.Commands[command].Pattern != pattern {
		t.Fatalf(`it should load dummy command with no pattern`)
	}
	if powershell.CommandMap.Commands[command].Description != description {
		t.Fatalf(`it should load dummy command with description`)
	}
}

func assertPromptName(powershell *Shell, name string, t *testing.T) {
	if powershell.Prompt != name {
		t.Fatalf(`it should be prompt name %s`, name)
	}
}

func assertShouldLoadSuggestion(powershell *Shell, text string, description string, t *testing.T) {
	suggestions := powershell.Suggestions
	if len(suggestions) != 1 {
		t.Fatalf(`it should load suggestion with size 1`)
	}
	if suggestions[0].Text != text {
		t.Fatalf(`it should load suggestions with test %s`, text)
	}
	if suggestions[0].Description != description {
		t.Fatalf(`it should load suggestions with test %s`, text)
	}
}

func assertShouldLoadCommandAction(powershell *Shell, command string, t *testing.T) {
	if powershell.CommandMap.Commands[command].Action.Function == nil {
		t.Fatalf(`it should load dummy command with action`)
	}
}

func assertShouldLoadCommandInputAction(powershell *Shell, command string, t *testing.T) {
	if powershell.CommandMap.Commands[command].Action.InputFunction == nil {
		t.Fatalf(`it should load dummy command with input action`)
	}
}

func assertShouldNotLoadCommandAction(powershell *Shell, command string, t *testing.T) {
	if powershell.CommandMap.Commands[command].Action.Function != nil {
		t.Fatalf(`it should not load dummy command with action`)
	}
}

func assertShouldNotLoadCommandInputAction(powershell *Shell, command string, t *testing.T) {
	if powershell.CommandMap.Commands[command].Action.InputFunction != nil {
		t.Fatalf(`it should load dummy command with input action`)
	}
}

func assertShouldNotLoadCommands(powershell *Shell, t *testing.T) {
	if powershell.CommandMap != nil {
		t.Fatalf(`it should not load commands`)
	}
}

func assertShouldNotLoadCommandsWithInvalidConfig(powershell *Shell, t *testing.T) {
	if len(powershell.CommandMap.Commands) > 0 {
		t.Fatalf(`it should not load commands because invalid config`)
	}
}

func assertShouldNotLoadSuggestions(powershell *Shell, t *testing.T) {
	if len(powershell.Suggestions) > 0 {
		t.Fatalf(`it should not load suggestions`)
	}
}

func assertShouldNotLoadActions(powershell *Shell, t *testing.T) {
	if powershell.ActionMap != nil {
		t.Fatalf(`it should not load actions`)
	}
}

func assertShouldNotLoadInputActions(powershell *Shell, t *testing.T) {
	if powershell.ActionOneStringMap != nil {
		t.Fatalf(`it should not load actions input`)
	}
}
