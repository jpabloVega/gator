package main

import (
	"fmt"
	"gator/internal/config"
	"os"
)

func main() {
	configData, err := config.ReadConfig()
	if err != nil {
		fmt.Println(err)
	}
	state := state{
		config: &configData,
	}
	cmds := StartCommands()
	cmds.register("login", handlerLogin)
	userArgs := os.Args
	if len(userArgs) < 2 {
		fmt.Println("Not enough arguments provided")
		os.Exit(1)
	}
	userCmd := GetUserCommand(userArgs)
	err = cmds.run(&state, userCmd)
	if err != nil {
		os.Exit(1)
	}
}
