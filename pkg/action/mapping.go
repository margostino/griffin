package action

var ActionStorage = map[string]func(){
	"ExecuteExit":        ExecuteExit,
	"ExecuteShowMessage": ExecuteShowMessage,
	"ExecuteHelp":        ExecuteHelp,
}

var InputActionStorage = map[string]func([]string){
	"ExecuteSelectInput": ExecuteSelectInput,
}
