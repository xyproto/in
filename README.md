![logo](img/in_160.png)

# in ![Build](https://github.com/xyproto/in/workflows/Build/badge.svg)

A utility for running a command within another directory (or directories matching a glob pattern).

It will also create the directories, if missing. If the top level directory is empty after executing the command, it will be removed. This means that `in testdirectory pwd` leaves no traces. When running in multiple directories, it will not create any new directories.

## Example 1

    in build cmake ..

Instead of:


    mkdir -p build
    cd build
    cmake ..

Or:

    mkdir -p build
    cmake -B build -S .

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

Globbing (note the double quotes to avoid shell expansion):

    in "./**/*pom.xml" mvn clean

## Installation

Either download the binary release (for 64-bit Linux), or install the development version (using `go install` like this requires Go 1.19 or later):

    go install github.com/xyproto/in@latest

Manual installation, using `git`, `go`, `sudo` and `install`:

    git clone https://github.com/xyproto/in
    cd in
    go build
    sudo install -Dm755 in /usr/bin/in

## Dependencies

* Go 1.19 or later

When compiling with GCC 10.2.0 (`gccgo`), the `in` executable is only 41k here.

## General info

* Version: 1.6.0
* License: BSD-3
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
