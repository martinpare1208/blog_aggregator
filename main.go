package main

import (
	"fmt"
	"log"
	"os"

	"github.com/martinpare1208/gator/internal/command"
	"github.com/martinpare1208/gator/internal/config"
)


func main() {
	data, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	fmt.Printf("Read config: %+v\n", data)
	
	state := &command.State{
		CfgPtr: &data, 
	}

	commands := command.Commands{
		CliCommands: make(map[string]func(*command.State, command.Command) error),
		}

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		return
	}

	commands.Register("login", command.HandlerLogin)


	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = commands.Run(state, command.Command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}


}