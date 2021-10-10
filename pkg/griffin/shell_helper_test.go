package griffin

import (
	"fmt"
)

var ActionMap = map[string]func(){
	"ExecuteDummyAction": ExecuteDummyAction,
}

var ActionOneString = map[string]func([]string){
	"ExecuteSelectInput": ExecuteDummyInputAction,
}

func ExecuteDummyAction() {
	fmt.Println("dummy action")
}

func ExecuteDummyInputAction(args []string) {
	fmt.Printf("dummy action with input %s\n", args[0])
}
