package griffin

import (
	"fmt"
	"os"
)

type Action struct {
	Function      func()
	InputFunction func([]string)
}

func NewAction(function func()) *Action {
	return &Action{
		Function: function,
	}
}

func NewInputAction(function func([]string)) *Action {
	return &Action{
		InputFunction: function,
	}
}

func (a Action) Execute() {
	a.Function()
}

// ExecuteExit TODO: define default actions custom properly
func ExecuteExit() {
	fmt.Println("bye!")
	os.Exit(0)
}
