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

Download the latest release and build it with `go build`, then install it with `install -Dm755 in /usr/bin/in`.

Or download the binary release (for 64-bit Linux).

Or, for the development release:

    export GOPATH=~/go
    go get github.com/xyproto/in
    install -Dm755 "$GOPATH/bin/in" /usr/bin/in

## Version

1.2
