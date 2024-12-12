package config

import (
	"fmt"
	"github.com/martinpare1208/gator/internal/read"
)

// "encoding/json"



type Config struct {
	DbURL string `json:"db_url"`
}

func Read() (error) {
	// var config Config
	// err := json.Unmarshal(jsonFile, &config)

	dir, err := read.GetConfigFilePath()
	if err != nil {
		return err
	}

	fmt.Println(dir)

	return nil

}


