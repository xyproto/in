package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

const versionString = "in 1.4.0"

// Enter a directory, creating the directory first if needed.
// Return true if a directory was created.
func enterAndCreate(path string) (bool, error) {
	err := os.Chdir(path)
	if err == nil { // success, no need to create the directory first
		return false, nil
	}
	// for any other error but "no such file or directory", return with an error
	if pe, ok := err.(*os.PathError); ok && pe.Err.Error() != "no such file or directory" {
		return false, err
	}
	// create the missing directory
	if err = os.MkdirAll(path, 0755); err != nil {
		return false, err
	}
	// enter the directory
	return true, os.Chdir(path)
}

// Run a command
func run(args ...string) error {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return err
	}
	return cmd.Wait()
}

func main() {
	if len(os.Args) <= 1 {
		log.Fatalln("too few arguments, need a directory as the first argument")
	}
	if os.Args[1] == "-v" || os.Args[1] == "--version" {
		fmt.Println(versionString)
		return
	}
	if len(os.Args) <= 2 {
		log.Fatalln("too few arguments, need a command after the first argument")
	}
	startDir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	dirName := os.Args[1]
	// enter the given directory (and create it, if needed)
	ok, err := enterAndCreate(dirName)
	if err != nil {
		log.Fatalln(err)
	}
	// run the given command
	if err := run(os.Args[2:]...); err != nil {
		if ok { // remove the created directory, if it's empty
			os.Remove(filepath.Join(startDir, dirName))
		}
		log.Fatalln(err) // exit(1) and skip the deferred function
	} else if ok { // remove the created directory, if it's empty
		os.Remove(filepath.Join(startDir, dirName))
	}
}
