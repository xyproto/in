package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const versionString = "in 1.2.0"

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Too few arguments, need a directory as the first argument")
		os.Exit(1)
	}

	if os.Args[1] == "-v" || os.Args[1] == "--version" {
		fmt.Println(versionString)
		os.Exit(0)
	}

	dir := os.Args[1]
	if err := os.Chdir(dir); err != nil {
		if !strings.Contains(err.Error(), "no such file or directory") {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		// Create the missing directory, then try to enter
		if err = os.MkdirAll(dir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		if err = os.Chdir(dir); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
	}

	if len(os.Args) <= 2 {
		fmt.Println("Too few arguments, need a command after the first argument")
		fmt.Fprintln(os.Stderr, "Too few arguments, need a command after the first argument")
		os.Exit(1)
	}

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	if err := cmd.Wait(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
