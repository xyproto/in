package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

const versionString = "in 1.3.0"

// enter a directory, creating the directory first if needed
func enter(path string) error {
	err := os.Chdir(path)
	if err == nil { // success, no need to create the directory first
		return nil
	}
	// check if the returned PathError has the message "no such file or directory"
	if pe, ok := err.(*os.PathError); ok && pe.Err.Error() == "no such file or directory" {
		return err
	}
	// create the missing directory
	if err = os.MkdirAll(path, 0755); err != nil {
		return err
	}
	// enter the directory
	if err = os.Chdir(path); err != nil {
		return err
	}
	return nil
}

// run a command
func run(args ...string) error {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return err
	}
	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}

// the main program
func program() error {
	if len(os.Args) <= 1 {
		return errors.New("too few arguments, need a directory as the first argument")
	}
	if os.Args[1] == "-v" || os.Args[1] == "--version" {
		fmt.Println(versionString)
		os.Exit(0)
	}
	if len(os.Args) <= 2 {
		return errors.New("too few arguments, need a command after the first argument")
	}
	// enter the given directory
	if err := enter(os.Args[1]); err != nil {
		return err
	}
	// run the given command
	if err := run(os.Args[2:]...); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := program(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
