package main

import (
	"example.com/echo/action"
	"github.com/margostino/griffin/pkg/griffin"
)

func main() {
	powershell := griffin.New().
		SetActions(action.ActionMap).
		SetActionsStrings(action.ActionOneString).
		LoadConfiguration("./config/commands.yml")
	powershell.Start()
}
