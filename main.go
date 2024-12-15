package main

import (
	"fmt"
	"os"

	"github.com/zipsonic/gator/internal/config"
)

type state struct {
	config config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	function map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.function[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	if functionToRun, ok := c.function[cmd.name]; ok {
		functionToRun(s, cmd)
	} else {
		return fmt.Errorf("command not found")
	}
	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("must supply a username")
	}

	err := s.config.SetUser(cmd.args[0])

	if err != nil {
		return err
	}

	fmt.Println("user has been set to: ", s.config.CurrentUserName)

	return nil
}

func main() {

	var configstate state
	err := config.Read(&configstate.config)
	if err != nil {
		fmt.Println("error reading config: ", err)
	}

	var functions commands
	var logincmd command
	logincmd.name = "login"
	logincmd.args = os.Args

	functions.register(logincmd.name, handlerLogin(&configstate, logincmd))

	err = config.Read(&cfg)
	if err != nil {
		fmt.Println("Error Reading Config: ", err)
	}

	fmt.Println("db_url : ", *cfg.DBUrl)
	fmt.Println("current_user : ", *cfg.CurrentUserName)
}
