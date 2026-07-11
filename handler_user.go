package main

import (
	"context"
	"fmt"
	"time"

	"github.com/cocuum/aggreGator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("Usage: register <name> - name is required") 
	} else if len(cmd.Args) > 1 {
		return fmt.Errorf("Usage: register <name> - only ONE name")
	}

	name := cmd.Args[0]

	user,err := s.db.CreateUser(
		context.Background(),
		database.CreateUserParams{
			ID:			uuid.New(),
			CreatedAt:	time.Now().UTC(),
			UpdatedAt:	time.Now().UTC(),
			Name:		name,
	})
	if err != nil {
		return fmt.Errorf("Could not create user: %w", err)
	}
	
	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("Could not create user: %w", err)
	}
	
	fmt.Printf("User %s was created successfully.\n", user.Name)
	printUserInfo(user)

	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("Usage: login <username> - username is required") 
	} else if len(cmd.Args) > 1 {
		return fmt.Errorf("Usage: login <username> - only ONE username")
	}

	userName := cmd.Args[0]

	_,err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		return fmt.Errorf("User %s does not exist", userName)
	}

	err = s.cfg.SetUser(userName)
	if err != nil {
		return fmt.Errorf("Could not create user: %w", err)
	}
	fmt.Printf("Username set to: %s\n", userName)

	return nil
}

func printUserInfo(user database.User) {
	fmt.Printf(" ++ ID:		%v\n", user.ID)
	fmt.Printf(" ++ Name:	%v\n", user.Name)
}