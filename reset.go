package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.ResetDb(context.Background())
	if err != nil {
		return fmt.Errorf("error resetting db, %v\n", err)
	}

	fmt.Println("db reset successfully!")

	return nil
}
