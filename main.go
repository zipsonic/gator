package main

import (
	"fmt"

	"github.com/zipsonic/gator/internal/config"
)

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
