package action

import (
	"fmt"
	"os"
)

type Exit struct {
	Action
}

func ExecuteExit() {
	fmt.Println("command exit")
	os.Exit(0)
}
