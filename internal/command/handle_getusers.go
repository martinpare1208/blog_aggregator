package command

import (
	"context"
	"fmt"
	"log"
)



func HandlerGetUsers(s *State, cmd Command) error {

	// get user records
	data, err := s.DBConnection.GetUsers(context.Background())
	if err != nil {
		log.Fatal(err)
		return err
	}

	for _, user := range data {
		if s.CfgPtr.CurrentUser == user.Name {
			fmt.Printf("%s (current)\n", user.Name)
		} else {
			fmt.Printf("%s", user.Name)
		}
	}

	fmt.Printf("Data retrieval complete\n")
	return nil

}