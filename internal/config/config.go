package config

import (
	"encoding/json"
	"errors"
	"os"
)

const FileName = "/.gatorconfig.json"

type Config struct {
	DbURL string `json:"db_url"`
	CurrentUser string `json:"current_user"`
}


func (c *Config) SetUser(userName string) error {
	c.CurrentUser = userName
	return write(*c)
}

func write(cfg Config) error {
	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}


func Read() (Config, error) {

	// Get current file path
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, errors.New("could not get file path")
	}

	// Read the json
	content, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, errors.New("could not read json file")
	}

	// Unpack json into go struct
	var payload Config
	err = json.Unmarshal(content, &payload)
	if err != nil {
		return Config{}, errors.New("error during Unmarshal()")
	}

	return payload, nil



}

func getConfigFilePath() (string, error) {

	// Get user's dir
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}


	// Combine user's dir to config json
	jsonFilePath := dir + FileName
	return jsonFilePath, nil
}