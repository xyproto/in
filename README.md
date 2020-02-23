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

## Full source code

Here's the full source code for the utility, **main.go**:

```go
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
```

## General info

* Version: 1.2
* License: MIT
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
