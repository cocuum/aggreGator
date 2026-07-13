package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/cocuum/aggreGator/internal/config"
	"github.com/cocuum/aggreGator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db	*database.Queries
	cfg *config.Config
}

func main() {
	// read the config file
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	//Open a connection to the database, and store it in the state struct
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatal("Cannot open connection to database:", err)
	}
	defer db.Close()
	
	dbQueries := database.New(db)
	if dbQueries == nil {
		log.Fatal("No Queries returned!!")
	}
	
	// Store the config & database queries in a new instance of the state struct
	programState := &state{
		db: dbQueries,
		cfg: &cfg,
	}

	// Initialize the commands struct with an empty map
	cmds := commands{
		registeredCMDS: make(map[string]func(*state, command) error),
	}

	//Register handler functions
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)

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