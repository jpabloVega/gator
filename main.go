package main

import (
	"database/sql"
	"fmt"
	"gator/internal/api"
	"gator/internal/config"
	"gator/internal/database"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	// Get config data
	configData, err := config.ReadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Get database
	dBData, err := sql.Open("postgres", configData.Db_url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Get client
	client := api.NewClient(time.Duration(5000))

	// Convert db data into querry
	db := database.New(dBData)

	// Create state
	state := state{
		client: &client,
		db:     db,
		config: &configData,
	}

	// Create commands struct
	cmds := StartCommands()
	cmds.register("login", handlerLogin)
	cmds.register("register", registerUser)
	cmds.register("reset", resetTable)
	cmds.register("users", getUsers)
	cmds.register("agg", aggregator)
	userArgs := os.Args
	if len(userArgs) < 2 {
		fmt.Println("Not enough arguments provided")
		os.Exit(1)
	}

	// Pass arguments
	userCmd := GetUserCommand(userArgs)

	// Run function
	err = cmds.run(&state, userCmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
