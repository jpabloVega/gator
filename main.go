package main

import (
	"fmt"
	"gator/internal/config"
)

func main() {
	configData, err := config.ReadConfig()
	if err != nil {
		fmt.Println(err)
	}
	state := config.State{
		Config: &configData,
	}
	cmds := config.StartCommands()
	cmds.register("login", config.HandlerLogin())
}
