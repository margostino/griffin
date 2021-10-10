package context

import (
	"github.com/margostino/griffin/pkg/command"
	"strings"
)

type ExecutionPlan struct {
	Plan    string
	Command *command.Command
}

func Prepare(plan string) *ExecutionPlan {
	// Potentially pre-processing
	return &ExecutionPlan{Plan: plan}
}

func (e *ExecutionPlan) With(command *command.Command) *ExecutionPlan {
	e.Command = command
	return e
}

func (e *ExecutionPlan) Execute() {
	if e.Command.Args > 0 {
		fields := strings.Fields(e.Plan)
		args := fields[len(fields)-e.Command.Args:]
		e.Command.ExecuteWith(args)
	} else {
		e.Command.Execute()
	}
}
