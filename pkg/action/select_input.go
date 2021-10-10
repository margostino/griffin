package action

import (
	"fmt"
)

func ExecuteSelectInput(args []string) {
	input := args[0]
	fmt.Printf("command select input %s\n", input)
}
