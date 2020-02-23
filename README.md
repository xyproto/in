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

## General info

* Version: 1.2
* License: MIT
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
