package main

import (
	"fmt"

	"github.com/zipsonic/gator/internal/config"
)

type state struct {
	config *config.Config
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

	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("Must supply a username")
	}

	err := s.config.SetUser(cmd.args[0])

	if err != nil {
		return err
	}

	fmt.Println("User has been set to: ", s.config.CurrentUserName)

	return nil
}

func main() {

	var cfg config.Config
	err := config.Read(&cfg)
	if err != nil {
		fmt.Println("Error Reading Config: ", err)
	}

	username := "rick"

	cfg.SetUser(username)

	err = config.Read(&cfg)
	if err != nil {
		fmt.Println("Error Reading Config: ", err)
	}

	fmt.Println("db_url : ", *cfg.DBUrl)
	fmt.Println("current_user : ", *cfg.CurrentUserName)
}
