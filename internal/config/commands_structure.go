package config

type command struct {
	name      string
	arguments []string
}

type commands struct {
	command map[string]func(*State, command) error
}

func StartCommands() commands {
	return commands{
		command: make(map[string]func(*State, command) error),
	}
}
