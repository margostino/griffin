package main

import (
	"example.com/echo/action"
	"github.com/margostino/griffin/pkg/griffin"
)

func main() {
	echo := griffin.New().
		SetActions(action.ActionMap).
		SetActionsStrings(action.ActionOneString).
		LoadConfiguration("./config/commands.yml")
	action.PowerEchoShell = echo
	echo.Start()
}
