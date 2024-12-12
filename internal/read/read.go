package read

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/martinpare1208/gator/internal/config"
)


func Read(file string) (config.Config, error) {

	// Get current file path
	filePath, err := getConfigFilePath(file)
	if err != nil {
		return config.Config{}, errors.New("could not get file path")
	}

	// Read the json
	content, err := os.ReadFile(filePath)
	if err != nil {
		return config.Config{}, errors.New("could not read json file")
	}

	// Unpack json into go struct
	var payload config.Config
	err = json.Unmarshal(content, &payload)
	if err != nil {
		return config.Config{}, errors.New("error during Unmarshal()")
	}

	// Print data from payload
	fmt.Printf("dbURL: %s\n", payload.DbURL)

	return payload, nil



}

func getConfigFilePath(cfgFile string) (string, error) {

	// Get user's dir
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}


	// Combine user's dir to config json
	jsonFilePath := dir + cfgFile
	return jsonFilePath, nil
}