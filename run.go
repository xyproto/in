package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// run a command within the given directory
func run(directory string, commandArgs []string, verbose bool) error {
	if verbose {
		fmt.Printf("run(%s, %v, %v)\n", directory, commandArgs, verbose)
	}

	commandName := commandArgs[0]

	cmd := exec.Command(commandName, commandArgs[1:]...)

	// Convert the given directory to an absolute path
	absDirectory, err := filepath.Abs(directory)
	if err != nil {
		return err
	}

	if verbose {
		fmt.Printf("absdir: %s\ncommandName: %s\n", absDirectory, commandName)
	}

	cmd.Dir = absDirectory

	// Append the PWD and INDIR environment variables
	cmd.Env = append(os.Environ(), "PWD="+absDirectory)
	cmd.Env = append(cmd.Env, "INDIR="+absDirectory)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return err
	}
	return cmd.Wait()
}

// runInAllMatching runs the given command in all the directories matching the given glob pattern
func runInAllMatching(pattern string, commandArgs []string, verbose bool) error {
	if verbose {
		fmt.Printf("runInAllMatching(%s, %v, %v)\n", pattern, commandArgs, verbose)
	}

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

		if err := run(dir, commandArgs, verbose); err != nil {
			return err
		}
	}
	return nil
}
