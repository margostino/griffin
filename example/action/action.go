package action

import (
	"fmt"
	"os"
)

var ActionMap = map[string]func(){
	"ExecuteExit":        ExecuteExit,
	"ExecuteShowMessage": ExecuteShowMessage,
	"ExecuteHelp":        ExecuteHelp,
}

var ActionOneString = map[string]func([]string){
	"ExecuteSelectInput": ExecuteSelectInput,
}

func ExecuteExit() {
	fmt.Println("echo shell: command exit")
	os.Exit(0)
}

func ExecuteHelp() {
	fmt.Println("echo shell: command help")
}

func ExecuteSelectInput(args []string) {
	input := args[0]
	fmt.Printf("echo shell: command select input %s\n", input)
}

func ExecuteShowMessage() {
	fmt.Println("echo shell: command show message")
}
