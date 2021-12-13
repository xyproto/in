![logo](img/in_160.png)

# in ![Build](https://github.com/xyproto/in/workflows/Build/badge.svg)

A command for running a command within another directory or directories matching a glob pattern.

It will also create the directories, if missing. If the top level directory is empty after executing the command, it will be removed. This means that `in testdirectory pwd` leaves no traces. When running in multiple directories, it will not create any new directories.

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

## Example 3

Support for globs via `filepath.Glob` which allows you to do something like (note the quotes which prevent a shell like bash from handling the globbing itself):

    in "./**/*pom.xml" mvn clean


## Installation

Either download the binary release (for 64-bit Linux), or install the development version (using `go install` like this requires Go 1.17 or later):

    go install github.com/xyproto/in@latest

Manual installation, using `git`, `go`, `sudo` and `install`:

    git clone https://github.com/xyproto/in
    cd in
    go build
    sudo install -Dm755 in /usr/bin/in

## Dependencies

* Go 1.3 or later

When compiling with GCC 10.2.0 (`gccgo`), the `in` executable is only 41k here.

## Full source code

Here's the full source code for the utility. There is always room for improvement, but it's relatively short and with no external dependencies:

```go
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
```

## General info

* Version: 1.4.0
* License: MIT
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
