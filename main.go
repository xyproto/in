package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const versionString = "in 1.6.0"

// Enter a directory, creating the directory first if needed.
// Return true if a directory was created.
func enterAndCreate(path string) (bool, error) {
	err := os.Chdir(path)
	if err == nil { // success, no need to create the directory first
		return false, nil
	}

	if !os.IsNotExist(err) {
		return false, err
	}

	// create the missing directory
	if err = os.MkdirAll(path, 0755); err != nil {
		return false, err
	}

	// enter the directory
	return true, os.Chdir(path)
}

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

func main() {
	if len(os.Args) <= 1 {
		log.Fatalln("too few arguments, need a directory or glob pattern as the first argument")
	}

	if os.Args[1] == "-v" || os.Args[1] == "--version" {
		fmt.Println(versionString)
		return
	}

	if len(os.Args) <= 2 {
		log.Fatalln("too few arguments, need a command after the first argument")
	}

	dirOrPattern := os.Args[1]

	// if the provided directory includes a glob pattern
	if strings.Contains(dirOrPattern, "*") {
		if err := runInAllMatching(dirOrPattern); err != nil {
			log.Fatalln(err)
		}
		return
	}

	// enter the given directory (and create it, if needed)
	dirCreated, err := enterAndCreate(dirOrPattern)
	if err != nil {
		log.Fatalln(err)
	}

	// run the given command
	if err := run(dirOrPattern, os.Args[2:]...); err != nil {
		log.Fatalln(err)
	}

	// remove the created directory if it's empty, and if it was created by this program
	if dirCreated {
		files, err := os.ReadDir(dirOrPattern)
		if err != nil || len(files) == 0 {
			os.Remove(dirOrPattern)
		}
	}
}

