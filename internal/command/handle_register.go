package command

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/martinpare1208/gator/internal/database"
)



func HandlerRegister(s *State, cmd Command) error {
	if len (cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	// create new user in db
	user, err := s.DBConnection.CreateUser(context.Background(), database.CreateUserParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: cmd.Args[0]})
	if err != nil {
		log.Fatal(err)
		return err
	}

	s.CfgPtr.SetUser(user.Name)
	fmt.Printf("User: '%s' successfully created!\n", user.Name)
	return nil

}