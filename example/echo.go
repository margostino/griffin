package main

import (
	"example.com/echo/action"
	"github.com/margostino/griffin/pkg/griffin"
)

func main() {
	echo := griffin.New().
		SetPrompt("echo").
		SetSimpleActions(action.ActionMap).
		SetMultiParamsActions(action.ActionOneString).
		LoadConfiguration("./config/commands.yml")
	action.EchoShell = echo
	echo.Start()
}
