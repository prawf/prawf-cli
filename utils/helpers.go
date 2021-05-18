package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// GetDir gets the current working directory
func GetDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return dir, nil
}

// AskForConfirmation asks the user for confirmation. A user must type in "yes" or "no" and then press enter. It has fuzzy matching, so "y", "Y", "yes", "YES", and "Yes" all count as confirmations. If the input is not recognized, it will ask again. The function does not return until it gets a valid response from the user.
func AskForConfirmation(s string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [y/n]? ", s)

		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		response = strings.ToLower(strings.TrimSpace(response))

		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		}
	}
}

func FileExists(filePath string) bool {
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func GetFilePath(configPath string, configName string) (string, string) {
	fileName := configName + "." + ConfigType
	filePath := filepath.Join(configPath, fileName)

	return filePath, fileName
}
