package utils

import "os"

// GetDir gets the current working directory
func GetDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return dir, nil
}
