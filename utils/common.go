package utils

import (
	"os"
	"path/filepath"
)

// GetCWD return current working directory
func GetCWD() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir, err
}
