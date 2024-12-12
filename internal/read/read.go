package read

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/martinpare1208/gator/internal/config"
)

const configFileName = "/blog_aggregator/.gatorconfig.json"


func Read() (error) {

	// Get current file path
	filePath, err := getConfigFilePath()
	if err != nil {
		return errors.New("could not get file path")
	}

	// Read the json
	content, err := os.ReadFile(filePath)
	if err != nil {
		return errors.New("could not read json file")
	}

	// Unpack json into go struct
	var payload config.Config
	err = json.Unmarshal(content, &payload)
	if err != nil {
		return errors.New("error during Unmarshal()")
	}

	// Print data from payload
	fmt.Printf("dbURL: %s\n", payload.DbURL)

	return nil



}

func getConfigFilePath() (string, error) {

	// Get user's dir
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	// Combine user's dir to project dir
	jsonFilePath := dir + configFileName
	return jsonFilePath, nil
}