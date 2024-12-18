package command

import (
	"context"
	"fmt"
	"log"
)

func HandlerLogin(s *State, cmd Command) error {
	if len (cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	// Check if user is in database
	userData, err := s.DBConnection.GetUser(context.Background(), name)
	if err != nil {
		log.Fatal("error: user not found")
	}

	
	err = s.CfgPtr.SetUser(userData.Name)
	if err != nil {
		return err
	}
	
	fmt.Println("User switched")
	return nil
}