# Griffin

Shell generator built on top of [GO Prompt](https://github.com/c-bata/go-prompt) (a library for building powerful
interactive prompts). Griffin allows to bind command and actions by configuration in order to reduce boilerplate code
and let developers focus on the action handling.

```go
package main

import (
	"fmt"
	"github.com/margostino/griffin/pkg/griffin"
)

var Actions = map[string]func(){
	"ExecuteDoSomething": ExecuteDoSomething,
}

func ExecuteDoSomething() {
	fmt.Println("do something")
}

func main() {
	powershell := griffin.New().
		SetActions(Actions).
		LoadConfiguration("commands.yml")
	powershell.Start()
}
```

### Commands

A shell can be created with Actions and Commands. A command might be: `run job`.

### Actions

An action is a function name without params: `ExecuteSomething()` or a function name with string
params: `ExecuteSomething(name string)`. Currently, only string params are supported.

### Configuration

Configuration might be loaded by YML file (see [example](./example/config/commands.yml)) or sending a list
of `CommandConfiguration`:

```go
type CommandConfiguration struct {
    Id          string `yaml:"id"`
    Description string `yaml:"description"`
    Args        int    `yaml:"args"`
    Action      string `yaml:"action"`
    Pattern     string `yaml:"pattern"`
}
```

`ID`: the command keys to trigger an action.  
Example: `run job`, `help`, `show users`

`Description`: short description of the command and action. This will be visible in as prompt suggestion.

`Args`: amount of parameters that the action needs to be executed. Default `0`

`Action`: name of the function handler to serve the command ID.

`Pattern`: if `Args > 1` a regex pattern is used to match the command with action.

