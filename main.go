package main

import (
	"log"
	"os"

	"github.com/cocuum/aggreGator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	// read the config file
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}
	
	// Store the config in a new instance of the state struct
	programState := &state{
		cfg: &cfg,
	}

	// Initialize the commands struct with an empty map
	cmds := commands{
		registeredCMDS: make(map[string]func(*state, command) error),
	}

	//Register a handler function for the login command
	cmds.register("login", handlerLogin)

	//Test and collect args from user input
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <commands> [args...]", os.Args[0])
	}
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	//Use the commands.run method to run the given command and print any errors returned
	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}

}