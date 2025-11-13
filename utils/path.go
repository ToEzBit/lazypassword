package utils

import (
	"os"
	"path/filepath"
)

func GetAppDirectoryPath() string {

	homeDir, err := os.UserHomeDir()

	if err != nil {
		return ""
	}

	return filepath.Join("./")
	return filepath.Join(homeDir, ".local", "share", "lazypassword")

}
