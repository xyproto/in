package main

import (
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

// createAndEnter will create a directory if needed, and then enter it.
// Returns the number of directories created.
func createAndEnter(path string) (int, error) {
	if err := os.Chdir(path); err == nil { // success
		return 0, nil
	} else if !os.IsNotExist(err) { // an error that is something else than a missing directory
		return 0, err
	}
	// count number of directories in path
	count := strings.Count(path, string(os.PathSeparator))
	// try to create the missing directories
	if err := os.MkdirAll(path, 0o755); err != nil {
		return 0, err
	}
	// enter the directory
	if err := os.Chdir(path); err != nil {
		return count, err
	}
	return count, nil
}

// removeIfEmpty will remove the directory structure if all directories are empty
func removeIfEmpty(path string, depth int) error {
	if depth <= 0 {
		return nil
	}
	err := os.Remove(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // it's already gone, which is fine
		}
		if pathErr, ok := err.(*os.PathError); ok && pathErr.Err == syscall.ENOTEMPTY {
			return nil // directory isn't empty, which is fine
		}
		return err
	}
	// if successful, move up to the parent directory and try again
	return removeIfEmpty(filepath.Dir(path), depth-1)
}
