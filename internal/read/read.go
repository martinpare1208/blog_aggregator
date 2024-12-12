package read

import (
	"os"
)

const configFileName = "/blog_aggregator/internal/config/.gatorconfig.json"

func GetConfigFilePath() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	jsonFilePath := dir + configFileName
	return jsonFilePath, nil
}