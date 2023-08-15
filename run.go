package main

import (
	"os"
	"os/exec"
	"path/filepath"
)

// Run a command within a directory
func run(directory string, args ...string) error {
	commandName := args[0]

	cmd := exec.Command(commandName, args[1:]...)
	cmd.Dir = directory

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return err
	}
	return cmd.Wait()
}

// runInAllMatching runs the given command in all the directories matching the given glob pattern
func runInAllMatching(pattern string) error {
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return err
	}

	seenDirs := map[string]bool{}
	for _, match := range matches {
		dir := filepath.Dir(match)
		if seenDirs[dir] {
			continue // skip directories we've already seen
		}
		seenDirs[dir] = true

		if err := run(dir, os.Args[2:]...); err != nil {
			return err
		}
	}
	return nil
}
