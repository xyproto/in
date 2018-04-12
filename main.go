package main

import (
	"fmt"
	//	"io/ioutil"
	"os"
	"os/exec"
)

const versionString = "in 1.0"

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
	err := os.Chdir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	if len(os.Args) <= 2 {
		fmt.Println("Too few arguments, need a command after the first argument")
		fmt.Fprintln(os.Stderr, "Too few arguments, need a command after the first argument")
		os.Exit(1)
	}

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

}
