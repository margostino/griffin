package main

var ActionMap = map[string]func(){
	"ExecuteExit":        ExecuteExit,
	"ExecuteShowMessage": ExecuteShowMessage,
	"ExecuteHelp":        ExecuteHelp,
}

var ActionOneString = map[string]func([]string){
	"ExecuteSelectInput": ExecuteSelectInput,
}
