package main

import "github.com/margostino/griffin/pkg/shell"

func main() {
	powershell := shell.NewShell("./config/commands.yml")
	powershell.Start()
}
