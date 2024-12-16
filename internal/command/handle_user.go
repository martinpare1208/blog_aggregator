package command

import (
	"fmt"
)

func HandlerLogin(s *State, cmd Command) error {
	if len (cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]
	err := s.CfgPtr.SetUser(name)
	if err != nil {
		return err
	}
	
	fmt.Println("User switched")
	return nil
}