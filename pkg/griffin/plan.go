package griffin

import (
	"strings"
)

type ExecutionPlan struct {
	Plan    string
	Command *Command
}

func Prepare(plan string) *ExecutionPlan {
	// Potentially pre-processing
	return &ExecutionPlan{Plan: plan}
}

func (e *ExecutionPlan) With(command *Command) *ExecutionPlan {
	e.Command = command
	return e
}

func (e *ExecutionPlan) Execute() {
	if e.Command.Args > 0 {
		args := strings.ReplaceAll(e.Plan, e.Command.Id, "")
		params := strings.Fields(args)
		e.Command.ExecuteWith(params)
	} else {
		e.Command.Execute()
	}
}
