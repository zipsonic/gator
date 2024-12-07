package config

import (
	"encoding/json"
	"os"
)

const CONFIGFILE string = ".gatorconfig.json"

type Config struct {
	DBUrl           *string `json:"db_url"`
	CurrentUserName *string `json:"current_user_name"`
}

func Read() (Config, error) {

	var config Config
	configpath, err := os.UserHomeDir()
	if err != nil {
		return config, err
	}
	configfile := configpath + "/" + CONFIGFILE

	file, err := os.Open(configfile)
	if err != nil {
		return config, err
	}

	defer file.Close()

	body := make([]byte, 0)

	_, err = file.Read(body)
	if err != nil {
		return config, err
	}

	if err = json.Unmarshal(body, &config); err != nil {
		println("Unmarshall error")
		return config, err
	}

	return config, nil
}
