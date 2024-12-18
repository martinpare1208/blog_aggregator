package command

import (
	"errors"

	"github.com/martinpare1208/gator/internal/config"
	"github.com/martinpare1208/gator/internal/database"
)

type State struct {
	CfgPtr *config.Config
	DBConnection *database.Queries
}

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	CliCommands map[string]func(*State, Command) error
}



func (c *Commands) Register(name string, f func(*State, Command) error) (error) {
	c.CliCommands[name] = f
	return nil
}

func (c *Commands) Run(s *State, cmd Command) (error) {

	f, ok := c.CliCommands[cmd.Name]
	if !ok {
		return errors.New("command not found")
	}

	return f(s, cmd)
}