package command

import (
	"context"
	"fmt"
	"log"
)



func HandlerReset(s *State, cmd Command) error {

	// reset db records
	err := s.DBConnection.Reset(context.Background())
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Printf("DB has been successfully reset\n")
	return nil

}