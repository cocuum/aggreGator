package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("Usage: login <username> - username is required") 
	} else if len(cmd.Args) > 1 {
		return fmt.Errorf("Usage: login <username> - only ONE username")
	}

	userName := cmd.Args[0]

	err := s.cfg.SetUser(userName)
	if err != nil {
		return fmt.Errorf("Username is not set correctly")
	}
	fmt.Printf("Username set to: %s\n", userName)

	return nil
}