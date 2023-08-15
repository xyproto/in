package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

// getParentPermissions retrieves the file permissions of the closest existing parent of the given directory.
func getParentPermissions(dir string) (os.FileMode, error) {
	parent := filepath.Dir(dir)
	_, err := os.Stat(parent)
	if os.IsNotExist(err) {
		// If the parent doesn't exist, try the parent's parent, and so on.
		return getParentPermissions(parent)
	} else if err != nil {
		return 0, err
	}

	// If parent exists, retrieve its permissions
	info, err := os.Stat(parent)
	if err != nil {
		return 0, err
	}
	return info.Mode().Perm(), nil
}

// createAndEnter will create a directory if needed, and then enter it.
// Returns the number of directories created.
func createAndEnter(path string, verbose bool) (int, error) {
	if verbose {
		fmt.Printf("createAndEnter(%s)\n", path)
	}

	if err := os.Chdir(path); err == nil { // success
		return 0, nil
	} else if !os.IsNotExist(err) { // an error that is something else than a missing directory
		return 0, err
	}

	perm, err := getParentPermissions(path)
	if err != nil {
		perm = 0o755 // default permission if parent permission couldn't be determined
	}

	// Split the path into parts
	parts := strings.Split(path, string(os.PathSeparator))
	if len(parts) == 0 {
		return 0, fmt.Errorf("invalid path")
	}

	dirCount := 0
	currentPath := parts[0]

	if len(path) > 0 && path[0] == os.PathSeparator {
		currentPath = string(os.PathSeparator) + currentPath
	}

	// Start from the root or the first directory and try creating each one
	for _, part := range parts[1:] {
		currentPath = filepath.Join(currentPath, part)
		if err := os.Mkdir(currentPath, perm); err != nil {
			if os.IsExist(err) {
				continue
			}
			return dirCount, err
		}
		dirCount++
	}

	// Enter the directory
	return dirCount, os.Chdir(path)
}

// removeIfEmpty will remove the directory structure if all directories are empty
func removeIfEmpty(path string, depth int, verbose bool) error {
	if verbose {
		fmt.Printf("removeIfEmpty(%s, %d)\n", path, depth)
	}

	if depth <= 0 {
		return nil
	}

	// Check if the path is a directory before attempting to remove it
	if info, err := os.Stat(path); err == nil && !info.IsDir() {
		return fmt.Errorf("path is not a directory: %s", path)
	}

	if err := os.Remove(path); err != nil {
		if os.IsNotExist(err) {
			return nil // it's already gone, which is fine
		}
		if pathErr, ok := err.(*os.PathError); ok && pathErr.Err == syscall.ENOTEMPTY {
			return nil // directory isn't empty, which is fine
		}
		return err
	}
	// if successful, move up to the parent directory and try again
	return removeIfEmpty(filepath.Dir(path), depth-1, verbose)
}
