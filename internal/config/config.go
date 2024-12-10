package config

import (
	"fmt"
	"os"
)

// "encoding/json"

type Config struct {
	DbURL string `json:"db_url"`
}

func Read() (error) {
	// var config Config
	// err := json.Unmarshal(jsonFile, &config)

	dir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	fmt.Println(dir)
	return nil

}

