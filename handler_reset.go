package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.Args) > 0 {
		return fmt.Errorf("Usage: reset - no arguments are required")
	}
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Could not reset users: %w", err)
	}
	fmt.Printf("Users was successfully reset.\n")

	return nil
}
