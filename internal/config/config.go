package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const CONFIGFILE string = ".gatorconfig.json"

type Config struct {
	DBUrl           *string `json:"db_url"`
	CurrentUserName *string `json:"current_user_name"`
}

func (c *Config) SetUser(username string) error {

	c.CurrentUserName = &username

	return Write(*c)
}

func Write(config Config) error {

	configpath, _ := getConfigPath()

	body, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling struct to JSON:", err)
		return err
	}

	file, err := os.Create(configpath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer file.Close()

	// Write the JSON data to the file
	_, err = file.Write(body)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}
	return nil
}

func getConfigPath() (string, error) {

	configpath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return (configpath + "/" + CONFIGFILE), nil
}

func Read(config *Config) error {

	configpath, err := getConfigPath()
	if err != nil {
		return err
	}

	file, err := os.Open(configpath)
	if err != nil {
		return err
	}

	defer file.Close()

	body, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, &config); err != nil {
		return err
	}

	return nil
}
