package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/martinpare1208/gator/internal/command"
	"github.com/martinpare1208/gator/internal/config"
	"github.com/martinpare1208/gator/internal/database"
)


func main() {
	data, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	fmt.Printf("Read config: %+v\n", data)
	
	
	// connect to database
	
	db, err := sql.Open("postgres", data.DbURL)
	if err != nil {
		log.Fatal(err)
	}
	
	dbQueries := database.New(db)
	
	state := &command.State{
		CfgPtr: &data,
		DBConnection: dbQueries, 
	}

	commands := command.Commands{
		CliCommands: make(map[string]func(*command.State, command.Command) error),
		}

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	commands.Register("login", command.HandlerLogin)
	commands.Register("register", command.HandlerRegister)
	commands.Register("reset", command.HandlerReset)
	commands.Register("users", command.HandlerGetUsers)
	
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = commands.Run(state, command.Command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}



}