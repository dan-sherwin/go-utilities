package utilities

import (
	"fmt"
	"os"
)

// AmAdmin checks if the current user has administrative (root) privileges by verifying if the effective user ID is 0. Returns true if the user is root; otherwise, false.
func AmAdmin() bool {
	return os.Geteuid() == 0
}

// DirCreateIfNotExists checks if a directory exists and creates it if it does not.
// It returns an error if there is an issue checking the directory or creating it.
func DirCreateIfNotExists(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return fmt.Errorf("Failed to create directory %s: %v", dir, err)
		}
	} else if err != nil {
		return fmt.Errorf("Failed to check directory %s: %v", dir, err)
	}
	return nil
}
