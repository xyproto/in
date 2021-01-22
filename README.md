# in [![Build Status](https://travis-ci.com/xyproto/in.svg?branch=master)](https://travis-ci.com/xyproto/in)

A command for running commands in other directories.

It will also create the directories, if missing.

## Example 1

    in build cmake ..

Instead of:

    mkdir -p build
    cd build
    cmake ..

## Example 2

    in project ./configure --prefix=/usr

Instead of:

    cd project
    ./configure --prefix=/usr
    cd ..

Or:

    (cd project; ./configure --prefix=/usr)

Or:

    pushd project
    ./configure --prefix=/usr
    popd

## Installation

Download the binary release (for 64-bit Linux), or install the development version:

    go get -u github.com/xyproto/in

Manual installation, using `git`, `go`, `sudo` and `install`:

    git clone https://github.com/xyproto/in
    cd in
    go build
    sudo install -Dm755 in /usr/bin/in

## Dependencies

* Go 1.3 or later

When compiling with GCC 10.2.0 (`gccgo`), the `in` executable is only 41k here.

## Full source code

Here's the full source code for the utility, **main.go**. There is always room for improvement, but it's pretty short:

```go
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
```

## General info

* Version: 1.3.0
* License: MIT
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
